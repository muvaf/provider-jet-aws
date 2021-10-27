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

type TagOptionResourceAssociationObservation struct {
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	ResourceCreatedTime *string `json:"resourceCreatedTime,omitempty" tf:"resource_created_time,omitempty"`

	ResourceDescription *string `json:"resourceDescription,omitempty" tf:"resource_description,omitempty"`

	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

type TagOptionResourceAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Required
	TagOptionID *string `json:"tagOptionId" tf:"tag_option_id,omitempty"`
}

// TagOptionResourceAssociationSpec defines the desired state of TagOptionResourceAssociation
type TagOptionResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagOptionResourceAssociationParameters `json:"forProvider"`
}

// TagOptionResourceAssociationStatus defines the observed state of TagOptionResourceAssociation.
type TagOptionResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagOptionResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagOptionResourceAssociation is the Schema for the TagOptionResourceAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TagOptionResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagOptionResourceAssociationSpec   `json:"spec"`
	Status            TagOptionResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagOptionResourceAssociationList contains a list of TagOptionResourceAssociations
type TagOptionResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagOptionResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	TagOptionResourceAssociationKind             = "TagOptionResourceAssociation"
	TagOptionResourceAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: TagOptionResourceAssociationKind}.String()
	TagOptionResourceAssociationKindAPIVersion   = TagOptionResourceAssociationKind + "." + GroupVersion.String()
	TagOptionResourceAssociationGroupVersionKind = GroupVersion.WithKind(TagOptionResourceAssociationKind)
)

func init() {
	SchemeBuilder.Register(&TagOptionResourceAssociation{}, &TagOptionResourceAssociationList{})
}