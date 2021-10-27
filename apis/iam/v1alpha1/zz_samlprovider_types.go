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

type SamlProviderObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type SamlProviderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SamlMetadataDocument *string `json:"samlMetadataDocument" tf:"saml_metadata_document,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SamlProviderSpec defines the desired state of SamlProvider
type SamlProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SamlProviderParameters `json:"forProvider"`
}

// SamlProviderStatus defines the observed state of SamlProvider.
type SamlProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SamlProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SamlProvider is the Schema for the SamlProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SamlProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SamlProviderSpec   `json:"spec"`
	Status            SamlProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SamlProviderList contains a list of SamlProviders
type SamlProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SamlProvider `json:"items"`
}

// Repository type metadata.
var (
	SamlProviderKind             = "SamlProvider"
	SamlProviderGroupKind        = schema.GroupKind{Group: Group, Kind: SamlProviderKind}.String()
	SamlProviderKindAPIVersion   = SamlProviderKind + "." + GroupVersion.String()
	SamlProviderGroupVersionKind = GroupVersion.WithKind(SamlProviderKind)
)

func init() {
	SchemeBuilder.Register(&SamlProvider{}, &SamlProviderList{})
}