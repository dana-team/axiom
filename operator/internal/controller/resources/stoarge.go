package resources

import (
	"context"

	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/go-logr/logr"
	storagev1 "k8s.io/api/storage/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetStorageProvisioners retrieves a list of storage provisioners from the cluster by listing all StorageClasses
// and extracting their provisioner information. This information is used to determine available storage options
// in the cluster.
func GetStorageProvisioners(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]v1alpha1.StorageProvisioner, error) {
	scList := &storagev1.StorageClassList{}
	if err := k8sClient.List(ctx, scList); err != nil {
		logger.Error(err, "Failed to list StorageClasses")
		return nil, err
	}
	provisioners := []v1alpha1.StorageProvisioner{}
	for _, sc := range scList.Items {
		sp := v1alpha1.StorageProvisioner{
			Name:        sc.Name,
			Provisioner: sc.Provisioner,
		}
		provisioners = append(provisioners, sp)
	}
	return provisioners, nil
}
