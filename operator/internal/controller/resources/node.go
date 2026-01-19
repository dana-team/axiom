package resources

import (
	"context"
	"fmt"

	"github.com/dana-team/axiom/operator/internal/controller/common"

	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClusterNodes retrieves all nodes from the Kubernetes cluster using the provided client
func GetClusterNodes(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]corev1.Node, error) {
	nodeList := &corev1.NodeList{}
	if err := k8sClient.List(ctx, nodeList); err != nil {
		logger.Error(err, "failed to list nodes")
		return nil, err
	}
	return nodeList.Items, nil
}

// CalculateClusterCompute calculates total compute resources (CPU, Memory, Storage, Pods, GPU) across all provided nodes
func CalculateClusterCompute(nodes []corev1.Node) v1alpha1.ClusterResources {
	totalCPU := resource.NewQuantity(0, resource.BinarySI)
	totalMemory := resource.NewQuantity(0, resource.BinarySI)
	totalPods := resource.NewQuantity(0, resource.DecimalSI)
	totalStorage := resource.NewQuantity(0, resource.BinarySI)
	var totalGPU int64

	for _, node := range nodes {
		if cpu := node.Status.Capacity.Cpu(); cpu != nil {
			totalCPU.Add(*cpu)
		}
		if memory := node.Status.Capacity.Memory(); memory != nil {
			totalMemory.Add(*memory)
		}
		if storage := node.Status.Capacity.StorageEphemeral(); storage != nil {
			totalStorage.Add(*storage)
		}
		if pods := node.Status.Allocatable.Pods(); pods != nil {
			totalPods.Add(*pods)
		}
		if gpuQuantity, ok := node.Status.Capacity[common.GpuLabel]; ok {
			totalGPU += gpuQuantity.Value()
		}
	}

	return v1alpha1.ClusterResources{
		CPU:     totalCPU.String(),
		Memory:  common.FormatGiB(totalMemory),
		Pods:    totalPods.String(),
		Storage: common.FormatGiB(totalStorage),
		GPU:     fmt.Sprintf("%d", totalGPU),
	}
}

// FormatNodesInfo transforms the node list into a simplified NodeInfo structure containing essential node information
func FormatNodesInfo(nodes []corev1.Node) []v1alpha1.NodeInfo {
	nodesInfo := []v1alpha1.NodeInfo{}
	for _, node := range nodes {
		nodeInfo := v1alpha1.NodeInfo{
			Name:           node.Name,
			InternalIP:     node.Status.Addresses[0].Address,
			Hostname:       node.Status.Addresses[1].Address,
			KubeletVersion: node.Status.NodeInfo.KubeletVersion,
			OSImage:        node.Status.NodeInfo.OperatingSystem,
		}
		nodesInfo = append(nodesInfo, nodeInfo)
	}
	return nodesInfo
}
