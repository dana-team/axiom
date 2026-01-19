package v1

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ConditionList []Condition

type Condition struct {
	Type               ConditionType          `json:"type"`
	Status             corev1.ConditionStatus `json:"status"`
	Reason             ConditionReason        `json:"reason,omitempty"`
	Message            string                 `json:"message,omitempty"`
	LastHeartbeatTime  metav1.Time            `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime metav1.Time            `json:"lastTransitionTime,omitempty"`
}

type ConditionType string

type ConditionReason string

func NewCondition(conditionType ConditionType, status corev1.ConditionStatus, reason ConditionReason, message string) Condition {
	now := metav1.Time{Time: time.Now()}
	condition := Condition{
		Type:               conditionType,
		Status:             status,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  now,
		LastTransitionTime: now,
	}
	return condition
}

func (conditions *ConditionList) Set(conditionType ConditionType, status corev1.ConditionStatus, reason ConditionReason, message string) {
	condition := conditions.Find(conditionType)

	// If there isn't condition we want to change, add new one
	if condition == nil {
		condition := NewCondition(conditionType, status, reason, message)
		*conditions = append(*conditions, condition)
		return
	}

	now := metav1.Time{Time: time.Now()}

	// If there is different status, reason or message update it
	if condition.Status != status || condition.Reason != reason || condition.Message != message {
		if condition.Status != status {
			condition.LastTransitionTime = now
		}
		condition.Status = status
		condition.Reason = reason
		condition.Message = message
	}
	condition.LastHeartbeatTime = now
}

func (conditions ConditionList) Find(conditionType ConditionType) *Condition {
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return &conditions[i]
		}
	}
	return nil
}

type RawState []byte

// State contains the namestatectl yaml [1] as string instead of golang struct
// so we don't need to be in sync with the schema.
//
// [1] https://github.com/nmstate/nmstate/blob/base/libnmstate/schemas/operational-state.yaml
// +kubebuilder:validation:Type=object
type State struct {
	Raw RawState `json:"-"`
}

func NewState(raw string) State {
	return State{Raw: RawState(raw)}
}

// This overrides the State type [1] so we can do a custom marshaling of
// nmstate yaml without the need to have golang code representing the
// nmstate schema
// +kubebuilder:validation:Type=object
// +kubebuilder:validation:Schemaless
// +kubebuilder:pruning:PreserveUnknownFields
// [1] https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (State) OpenAPISchemaType() []string { return []string{"object"} }

// NodeNetworkConfigurationPolicySpec defines the desired state of NodeNetworkConfigurationPolicy
type NodeNetworkConfigurationPolicySpec struct {
	// NodeSelector is a selector which must be true for the policy to be applied to the node.
	// Selector which must match a node's labels for the policy to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Capture contains expressions with an associated name than can be referenced
	// at the DesiredState.
	// +optional
	Capture map[string]string `json:"capture,omitempty"`

	// +kubebuilder:validation:XPreserveUnknownFields
	// The desired configuration of the policy
	DesiredState runtime.RawExtension `json:"desiredState,omitempty"`

	// MaxUnavailable specifies percentage or number
	// of machines that can be updating at a time. Default is "50%".
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
}

// NodeNetworkConfigurationPolicyStatus defines the observed state of NodeNetworkConfigurationPolicy
type NodeNetworkConfigurationPolicyStatus struct {
	Conditions ConditionList `json:"conditions,omitempty" optional:"true"`

	// UnavailableNodeCount represents the total number of potentially unavailable nodes that are
	// processing a NodeNetworkConfigurationPolicy
	// +optional
	UnavailableNodeCount int `json:"unavailableNodeCount,omitempty" optional:"true"`
	// LastUnavailableNodeCountUpdate is time of the last UnavailableNodeCount update
	// +optional
	LastUnavailableNodeCountUpdate *metav1.Time `json:"lastUnavailableNodeCountUpdate,omitempty" optional:"true"`
}

const (
	NodeNetworkConfigurationPolicyConditionAvailable   ConditionType = "Available"
	NodeNetworkConfigurationPolicyConditionDegraded    ConditionType = "Degraded"
	NodeNetworkConfigurationPolicyConditionProgressing ConditionType = "Progressing"
)

var NodeNetworkConfigurationPolicyConditionTypes = [...]ConditionType{
	NodeNetworkConfigurationPolicyConditionAvailable,
	NodeNetworkConfigurationPolicyConditionDegraded,
	NodeNetworkConfigurationPolicyConditionProgressing,
}

const (
	NodeNetworkConfigurationPolicyConditionFailedToConfigure           ConditionReason = "FailedToConfigure"
	NodeNetworkConfigurationPolicyConditionSuccessfullyConfigured      ConditionReason = "SuccessfullyConfigured"
	NodeNetworkConfigurationPolicyConditionConfigurationProgressing    ConditionReason = "ConfigurationProgressing"
	NodeNetworkConfigurationPolicyConditionConfigurationNoMatchingNode ConditionReason = "NoMatchingNode"
)

// +kubebuilder:object:root=true

// NodeNetworkConfigurationPolicyList contains a list of NodeNetworkConfigurationPolicy
type NodeNetworkConfigurationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkConfigurationPolicy `json:"items"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.status==\"True\")].type",description="Status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.status==\"True\")].reason",description="Reason"
// +kubebuilder:resource:shortName=nncp,scope=Cluster
// +kubebuilder:storageversion

// NodeNetworkConfigurationPolicy is the Schema for the nodenetworkconfigurationpolicies API
type NodeNetworkConfigurationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeNetworkConfigurationPolicySpec   `json:"spec,omitempty"`
	Status NodeNetworkConfigurationPolicyStatus `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NodeNetworkConfigurationPolicy{}, &NodeNetworkConfigurationPolicyList{})
}
