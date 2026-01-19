/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"github.com/dana-team/axiom/operator/internal/db"

	axiomv1alpha1 "github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/dana-team/axiom/operator/internal/controller/status"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ClusterInfoReconciler reconciles a ClusterInfo object
type ClusterInfoReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=axiom.dana.io,resources=clusterinfo,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=axiom.dana.io,resources=clusterinfo/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=axiom.dana.io,resources=clusterinfo/finalizers,verbs=update
// +kubebuilder:rbac:groups=axiom.dana.io,resources=clusterversions,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups="",resources=pods/log,verbs=get;list;watch;create
// +kubebuilder:rbac:groups=route.openshift.io,resources=routes,verbs=get;list;watch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch
// +kubebuilder:rbac:groups=config.openshift.io,resources=clusterversions,verbs=get;list;watch
// +kubebuilder:rbac:groups=config.openshift.io,resources=oauths,verbs=get;list;watch
// +kubebuilder:rbac:groups=nmstate.io,resources=nodenetworkconfigurationpolicies,verbs=get;list;watch
// +kubebuilder:rbac:groups=storage.k8s.io,resources=storageclasses,verbs=get;list;watch
// +kubebuilder:rbac:groups=nmstate.io,resources=nodenetworkstates,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ClusterInfoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling ClusterInfo")

	clusterInfo := &axiomv1alpha1.ClusterInfo{}
	if err := r.Get(ctx, req.NamespacedName, clusterInfo); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := status.UpdateClusterInfoStatus(ctx, logger, *clusterInfo, r.Client); err != nil {
		return ctrl.Result{}, fmt.Errorf("Failed to update ClusterInfo status %s", err.Error())
	}
	logger.Info("ClusterInfo status updated successfully")

	go func(clusterInfo axiomv1alpha1.ClusterInfo) {
		db.InsertClusterInfoToMongo(logger, clusterInfo)
	}(*clusterInfo)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterInfoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&axiomv1alpha1.ClusterInfo{}).
		Named("clusterinfo").
		Complete(r)
}
