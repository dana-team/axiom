package resources

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/dana-team/axiom/operator/internal/controller/common"
	"github.com/go-logr/logr"
	nmstatev1 "github.com/nmstate/kubernetes-nmstate/api/v1beta1"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ApiPrefix = "/api/ipam/prefixes/"

// GetClusterSegments retrieves network segments for a cluster based on its configuration and node network state.
// If the cluster is hosted, segments are fetched from an external NetBox API. Otherwise, segments are derived from node configurations.
// Returns a list of unique segments or an error in case of failure.
func GetClusterSegments(ctx context.Context, logger logr.Logger, k8sClient client.Client, ci *v1alpha1.ClusterInfo, nodes []corev1.Node, clusterName string) ([]string, error) {
	if ci.Spec.HostedCluster {
		return getSegmentsFromNetBox(ctx, logger, clusterName)
	}

	return getSegmentsFromNodeNetworkState(ctx, logger, k8sClient, nodes)
}

// getSegmentsFromNetBox retrieves unique network segments from the NetBox API for a specified cluster.
// Returns a list of segments or an error if the retrieval or parsing of data fails.
func getSegmentsFromNetBox(ctx context.Context, logger logr.Logger, clusterName string) ([]string, error) {
	NetBoxURL := os.Getenv("NETBOX_URL")
	NetBoxToken := os.Getenv("NETBOX_TOKEN")
	if NetBoxURL == "" || NetBoxToken == "" {
		logger.Info("NETBOX_URL or NETBOX_TOKEN environment variables are not set. Skipping NetBox segments retrieval.")
		return nil, nil
	}

	baseURL, err := url.Parse(NetBoxURL + ApiPrefix)
	if err != nil {
		logger.Error(err, "Failed to parse NetBox URL")
		return nil, err
	}

	cluster := common.StripDomain(clusterName)

	params := url.Values{}
	params.Add("cf_Cluster", cluster)
	params.Add("limit", "0")
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL.String(), nil)
	if err != nil {
		logger.Error(err, "Failed to create NetBox request")
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", NetBoxToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	resp, err := c.Do(req)
	if err != nil {
		logger.Error(err, "Failed to send NetBox request")
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logger.Info("Failed to close NetBox response body")
		}
	}()

	if resp.StatusCode != http.StatusOK {
		logger.Error(fmt.Errorf("NetBox request failed with status code %d", resp.StatusCode), "Failed to get segments from NetBox")
		return nil, fmt.Errorf("NetBox request failed with status code %d", resp.StatusCode)
	}

	var results NetBoxResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed to parse NetBox response: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d prefixes for cluster %s", results.Count, cluster))

	var segments []string
	for _, prefix := range results.Results {
		segments = append(segments, prefix.Prefix)
	}

	return common.FilterUniqueStrings(segments), nil
}

// getSegmentsFromNodeNetworkState retrieves unique network segments from the NodeNetworkState of each provided node.
// Queries the Kubernetes API to fetch NodeNetworkState objects for the given nodes and parses associated IPv4 information.
// Returns a slice of unique network segments or an error if any operation fails.
func getSegmentsFromNodeNetworkState(ctx context.Context, logger logr.Logger, k8sClient client.Client, nodes []corev1.Node) ([]string, error) {
	var segments []string
	for _, node := range nodes {
		nns := &nmstatev1.NodeNetworkState{}
		if err := k8sClient.Get(ctx, client.ObjectKey{Name: node.Name}, nns); err != nil {
			if errors.IsNotFound(err) {
				logger.Info("Node has no NodeNetworkState", "node", node.Name)
				continue
			}
			logger.Error(err, "Failed to get NodeNetworkState", "node", node.Name)
			return nil, err
		}

		if len(nns.Status.CurrentState.Raw) == 0 {
			logger.V(1).Info("Empty CurrentState", "node", node.Name)
			continue
		}

		var state NodeNetworkStateCurrentState
		if err := yaml.Unmarshal(nns.Status.CurrentState.Raw, &state); err != nil {
			logger.Error(err, "Failed to unmarshal NodeNetworkState", "node", node.Name)
			return nil, err
		}

		logger.V(1).Info("NNS parsed", "node", node.Name, "interfaces", len(state.Interfaces))

		for _, iface := range state.Interfaces {
			for _, addr := range iface.Ipv4.Address {
				if addr.IP != "" && addr.PrefixLength > 0 {
					s, err := common.CreateSegmentFromIPAndPrefix(addr.IP, addr.PrefixLength)
					if err != nil {
						logger.Error(err, "Failed to create segment", "ip", addr.IP, "prefix", addr.PrefixLength)
						return nil, err
					}
					segments = append(segments, s)
				}
			}
		}
	}

	unique := common.FilterUniqueStrings(segments)
	logger.Info("Segments collected", "count", len(unique), "nodes", len(nodes))
	return unique, nil
}
