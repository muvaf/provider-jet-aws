/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GatewayUsagePlanKeyObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GatewayUsagePlanKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`

	// +kubebuilder:validation:Required
	KeyType *string `json:"keyType" tf:"key_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	UsagePlanID *string `json:"usagePlanId" tf:"usage_plan_id,omitempty"`
}

// GatewayUsagePlanKeySpec defines the desired state of GatewayUsagePlanKey
type GatewayUsagePlanKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayUsagePlanKeyParameters `json:"forProvider"`
}

// GatewayUsagePlanKeyStatus defines the observed state of GatewayUsagePlanKey.
type GatewayUsagePlanKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayUsagePlanKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayUsagePlanKey is the Schema for the GatewayUsagePlanKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GatewayUsagePlanKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayUsagePlanKeySpec   `json:"spec"`
	Status            GatewayUsagePlanKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayUsagePlanKeyList contains a list of GatewayUsagePlanKeys
type GatewayUsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayUsagePlanKey `json:"items"`
}

// Repository type metadata.
var (
	GatewayUsagePlanKeyKind             = "GatewayUsagePlanKey"
	GatewayUsagePlanKeyGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayUsagePlanKeyKind}.String()
	GatewayUsagePlanKeyKindAPIVersion   = GatewayUsagePlanKeyKind + "." + GroupVersion.String()
	GatewayUsagePlanKeyGroupVersionKind = GroupVersion.WithKind(GatewayUsagePlanKeyKind)
)

func init() {
	SchemeBuilder.Register(&GatewayUsagePlanKey{}, &GatewayUsagePlanKeyList{})
}