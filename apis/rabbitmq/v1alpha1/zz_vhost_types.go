/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VhostInitParameters struct {
}

type VhostObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the vhost.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VhostParameters struct {

	// The name of the vhost.
	// +crossplane:generate:reference:type=Name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Name to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Name to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

// VhostSpec defines the desired state of Vhost
type VhostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VhostParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VhostInitParameters `json:"initProvider,omitempty"`
}

// VhostStatus defines the observed state of Vhost.
type VhostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VhostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vhost is the Schema for the Vhosts API. Creates and manages a vhost on a RabbitMQ server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rabbitmq}
type Vhost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VhostSpec   `json:"spec"`
	Status            VhostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VhostList contains a list of Vhosts
type VhostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vhost `json:"items"`
}

// Repository type metadata.
var (
	Vhost_Kind             = "Vhost"
	Vhost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vhost_Kind}.String()
	Vhost_KindAPIVersion   = Vhost_Kind + "." + CRDGroupVersion.String()
	Vhost_GroupVersionKind = CRDGroupVersion.WithKind(Vhost_Kind)
)

func init() {
	SchemeBuilder.Register(&Vhost{}, &VhostList{})
}
