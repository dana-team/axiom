package resources

import (
	"context"

	"github.com/go-logr/logr"
	addmissionv1 "k8s.io/api/admissionregistration/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetMutatingWebhooks retrieves a list of all MutatingWebhookConfiguration names in the cluster
func GetMutatingWebhooks(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	mutatingWebhooks := &addmissionv1.MutatingWebhookConfigurationList{}
	if err := k8sClient.List(ctx, mutatingWebhooks); err != nil {
		logger.Error(err, "Failed to list MutatingWebhookConfiguration")
		return nil, err
	}
	mWebhooks := []string{}
	for _, wh := range mutatingWebhooks.Items {
		mWebhooks = append(mWebhooks, wh.Name)
	}
	return mWebhooks, nil
}

// GetValidatingWebhooks retrieves a list of all ValidatingWebhookConfiguration names in the cluster
func GetValidatingWebhooks(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	validatingWebhooks := addmissionv1.ValidatingWebhookConfigurationList{}
	if err := k8sClient.List(ctx, &validatingWebhooks); err != nil {
		logger.Error(err, "Failed to list ValidatingWebhookConfiguration")
		return nil, err
	}
	vWebhooks := []string{}
	for _, wh := range validatingWebhooks.Items {
		vWebhooks = append(vWebhooks, wh.Name)
	}
	return vWebhooks, nil
}
