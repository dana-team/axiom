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

package v1alpha1

import (
	"sort"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeInfo holds information about a node
type NodeInfo struct {
	Name           string `json:"name,omitempty"`
	InternalIP     string `json:"internalIP,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	OSImage        string `json:"osImage,omitempty"`
	KubeletVersion string `json:"kubeletVersion,omitempty"`
}

// ClusterResources describes resource capacity of the cluster.
type ClusterResources struct {
	CPU     string `json:"cpu,omitempty"`
	Memory  string `json:"memory,omitempty"`
	Pods    string `json:"pods,omitempty"`
	Storage string `json:"storage,omitempty"`
	GPU     string `json:"gpu,omitempty"`
}

type StorageProvisioner struct {
	Name        string `json:"name"`
	Provisioner string `json:"provisioner"`
}

type ClusterDnsConfig struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Nullable
	SearchDomains []string `json:"searchDomains,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Nullable
	Servers []string `json:"servers,omitempty"`
}

// ClusterInfoSpec defines the desired state of ClusterInfo.
type ClusterInfoSpec struct {
	HostedCluster bool `json:"hostedCluster,omitempty" bson:"hostedCluster,omitempty"`
}

type ClusterInfoStatus struct {
	Name                string               `json:"name,omitempty" bson:"name,omitempty"`
	ClusterID           string               `json:"clusterID,omitempty" bson:"clusterID,omitempty"`
	KubernetesVersion   string               `json:"kubernetesVersion,omitempty" bson:"kubernetesVersion,omitempty"`
	ClusterDnsConfig    ClusterDnsConfig     `json:"clusterDnsConfig,omitempty" bson:"clusterDnsConfig,omitempty"`
	ClusterResources    ClusterResources     `json:"clusterResources,omitempty" bson:"clusterResources,omitempty"`
	NodeInfo            []NodeInfo           `json:"nodeInfo,omitempty" bson:"nodeInfo,omitempty"`
	RouterLBAddresses   []string             `json:"routerLBAddress,omitempty" bson:"routerLBAddress,omitempty"`
	ApiServerAddresses  []string             `json:"apiServerAddresses,omitempty" bson:"apiServerAddresses,omitempty"`
	IdentityProviders   []string             `json:"identityProviders,omitempty" bson:"identityProviders,omitempty"`
	StorageProvisioners []StorageProvisioner `json:"storageProvisioners,omitempty" bson:"storageProvisioners,omitempty"`
	MutatingWebhooks    []string             `json:"mutatingWebhooks,omitempty" bson:"mutatingWebhooks,omitempty"`
	ValidatingWebhooks  []string             `json:"validatingWebhooks,omitempty" bson:"validatingWebhooks,omitempty"`
	Segments            []string             `json:"segments,omitempty" bson:"segments,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clusterinfo,singular=clusterinfo,scope=Cluster,shortName=ci

// ClusterInfo is the Schema for the clusterinfoes API.
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterInfoSpec   `json:"spec,omitempty"`
	Status ClusterInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInfoList contains a list of ClusterInfo.
type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterInfo{}, &ClusterInfoList{})
}

func (s *ClusterInfoStatus) Normalize() {
	sort.Strings(s.RouterLBAddresses)
	sort.Strings(s.ApiServerAddresses)
	sort.Strings(s.IdentityProviders)
	sort.Strings(s.MutatingWebhooks)
	sort.Strings(s.ValidatingWebhooks)
	sort.Strings(s.Segments)

	sort.Slice(s.NodeInfo, func(i, j int) bool {
		return s.NodeInfo[i].Name < s.NodeInfo[j].Name
	})

	sort.Slice(s.StorageProvisioners, func(i, j int) bool {
		return s.StorageProvisioners[i].Name < s.StorageProvisioners[j].Name
	})
}
