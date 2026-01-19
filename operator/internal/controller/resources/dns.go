package resources

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"

	nmstatev1 "github.com/dana-team/axiom/operator/api/nmstate/v1"
	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/go-logr/logr"
	yamlv3 "gopkg.in/yaml.v3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ReslovConfDir      = "/etc/resolv.conf"
	HostMountPath      = "/host/resolv.conf"
	ServiceAccountName = "axiom-operator-controller-manager"
)

// DnsResolverConfig represents the DNS resolver configuration structure
// that matches the YAML format used in NodeNetworkConfigurationPolicy
type DnsResolverConfig struct {
	Config struct {
		Server []string `yaml:"server"`
		Search []string `yaml:"search"`
	} `yaml:"config"`
}

// GetClusterDnsConfiguration retrieves DNS configuration from NodeNetworkConfigurationPolicy
// and converts it to ClusterDnsConfig format
func GetClusterDnsConfiguration(ctx context.Context, logger logr.Logger, k8sClient client.Client, ci *v1alpha1.ClusterInfo) (v1alpha1.ClusterDnsConfig, error) {
	if ci.Spec.HostedCluster {
		dnsConfig, err := getDNSFromResolveConf(ctx, k8sClient, logger)
		if err != nil {
			return v1alpha1.ClusterDnsConfig{}, err
		}
		return dnsConfig, nil
	}

	dnsConfig, err := getDNSFromNNCP(ctx, logger, k8sClient)
	if err != nil {
		return v1alpha1.ClusterDnsConfig{}, err
	}
	return dnsConfig, nil
}

func getDNSFromNNCP(ctx context.Context, logger logr.Logger, k8sClient client.Client) (v1alpha1.ClusterDnsConfig, error) {
	var state struct {
		DNSResolver *DnsResolverConfig `yaml:"dns-resolver"`
	}

	nncp := &nmstatev1.NodeNetworkConfigurationPolicy{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "node-resolver"}, nncp); err != nil {
		logger.Error(err, "failed to get NodeNetworkConfigurationPolicy")
		return v1alpha1.ClusterDnsConfig{}, err
	}

	if err := yamlv3.Unmarshal(nncp.Spec.DesiredState.Raw, &state); err != nil {
		logger.Error(err, "failed to unmarshal DesiredState")
	}

	return v1alpha1.ClusterDnsConfig{
			SearchDomains: state.DNSResolver.Config.Search,
			Servers:       state.DNSResolver.Config.Server,
		},
		nil
}

func getDNSFromResolveConf(ctx context.Context, k8sClient client.Client, logger logr.Logger) (v1alpha1.ClusterDnsConfig, error) {
	pod := newDNSReaderPod()
	existingPod := &corev1.Pod{}
	err := k8sClient.Get(ctx, client.ObjectKey{Name: pod.Name, Namespace: pod.Namespace}, existingPod)
	if err == nil {
		pod = existingPod
	} else if client.IgnoreNotFound(err) == nil {
		if err := k8sClient.Create(ctx, pod); err != nil {
			return v1alpha1.ClusterDnsConfig{}, err
		}
	} else {
		return v1alpha1.ClusterDnsConfig{}, err
	}

	defer func() {
		err = k8sClient.Delete(ctx, pod)
		if err != nil {
			logger.Info(fmt.Sprintf("failed to delete Pod %s/%s due to %v", pod.Namespace, pod.Name, err))
		}
	}()

	err = wait.PollUntilContextTimeout(
		ctx,
		1*time.Second,
		30*time.Second,
		true,
		func(ctx context.Context) (bool, error) {
			var p corev1.Pod
			if err := k8sClient.Get(ctx, client.ObjectKey{Name: pod.Name, Namespace: pod.Namespace}, &p); err != nil {
				return false, client.IgnoreNotFound(err)
			}

			if p.Status.Phase == corev1.PodSucceeded || p.Status.Phase == corev1.PodFailed {
				return true, nil
			}
			return false, nil
		},
	)

	if err != nil {
		return v1alpha1.ClusterDnsConfig{}, fmt.Errorf("timed out waiting for dns-reader pod: %w", err)
	}

	restCfg := ctrl.GetConfigOrDie()
	clientSet, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return v1alpha1.ClusterDnsConfig{}, err
	}

	logReq := clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &corev1.PodLogOptions{})
	stream, err := logReq.Stream(ctx)
	if err != nil {
		return v1alpha1.ClusterDnsConfig{}, err
	}
	defer func() {
		if cerr := stream.Close(); cerr != nil {
			logger.Error(cerr, "failed to close stream")
			return
		}
	}()

	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, stream)
	content := buf.String()
	return parseResolveConf(content), nil
}

func newDNSReaderPod() *corev1.Pod {
	runAsUser := int64(0)
	hostPathType := corev1.HostPathFile

	return &corev1.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name:      "dns-reader",
			Namespace: os.Getenv("POD_NAMESPACE"),
		},
		Spec: corev1.PodSpec{
			ServiceAccountName: ServiceAccountName,
			RestartPolicy:      corev1.RestartPolicyNever,
			HostNetwork:        true,
			Containers: []corev1.Container{
				{
					Name:  "reader",
					Image: os.Getenv("DNS_READER_IMAGE"),
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "host-resolv",
							MountPath: HostMountPath,
							ReadOnly:  true,
						},
					},
					Command: []string{"cat", HostMountPath},
					SecurityContext: &corev1.SecurityContext{
						RunAsUser: &runAsUser,
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "host-resolv",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: ReslovConfDir,
							Type: &hostPathType,
						},
					},
				},
			},
		},
	}
}

func parseResolveConf(content string) v1alpha1.ClusterDnsConfig {
	var servers []string
	var search []string

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				servers = append(servers, parts[1])
			}
		}

		if strings.HasPrefix(line, "search") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				search = append(search, parts[1:]...)
			}
		}
	}

	cfg := v1alpha1.ClusterDnsConfig{}
	cfg.Servers = servers

	if len(search) > 0 {
		cfg.SearchDomains = search
	} else {
		cfg.SearchDomains = []string{"cluster.local"}
	}

	return cfg
}
