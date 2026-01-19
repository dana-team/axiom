package resources

import (
	"context"

	"github.com/go-logr/logr"
	configv1 "github.com/openshift/api/config/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClusterVersionAndID queries the ClusterVersion resource to obtain cluster information.
// It retrieves the desired OpenShift version from Status and cluster ID from Spec fields.
func GetClusterVersionAndID(ctx context.Context, logger logr.Logger, k8sClient client.Client) (string, string, error) {
	cv := &configv1.ClusterVersion{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "version"}, cv); err != nil {
		logger.Error(err, "Failed to get ClusterVersion")
		return "", "", err
	}
	ocpVersion := cv.Status.Desired.Version
	clusterID := cv.Spec.ClusterID
	return ocpVersion, string(clusterID), nil
}
