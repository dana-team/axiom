package resources

import (
	"context"
	"net"
	"strings"

	"github.com/dana-team/axiom/operator/internal/controller/common"

	"github.com/go-logr/logr"
	v1 "github.com/openshift/api/route/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetRouterLBAddress retrieves the load balancer IP addresses for the OpenShift console route
// by looking up the DNS host of the console route
func GetRouterLBAddress(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	route := &v1.Route{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: common.ConsoleName, Namespace: common.ConsoleNamespace}, route); err != nil {
		logger.Error(err, "Failed to get console route")
		return nil, err
	}
	ips, err := net.LookupHost(route.Spec.Host)
	if err != nil {
		logger.Error(err, "Failed to lookup host")
		return nil, err
	}
	return ips, nil
}

// GetApiServerAddress retrieves the API server IP addresses by getting the console route
// and replacing its prefix with "api." to construct the API server URL
func GetApiServerAddress(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	route := &v1.Route{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: common.ConsoleName, Namespace: common.ConsoleNamespace}, route); err != nil {
		logger.Error(err, "Failed to get console route")
		return nil, err
	}
	apiServerURL := strings.Replace(route.Spec.Host, common.IngressPrefix, "api.", 1)
	ips, err := net.LookupHost(apiServerURL)
	if err != nil {
		logger.Error(err, "Failed to lookup host")
		return nil, err
	}
	return ips, nil
}

func GetClusterName(ctx context.Context, logger logr.Logger, k8sClient client.Client) (string, error) {
	route := &v1.Route{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: common.ConsoleName, Namespace: common.ConsoleNamespace}, route); err != nil {
		logger.Error(err, "Failed to get console route")
		return "", err
	}
	clusterName := strings.Replace(route.Spec.Host, common.IngressPrefix, "", 1)
	return clusterName, nil
}
