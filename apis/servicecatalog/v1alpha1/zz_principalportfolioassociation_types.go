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

type PrincipalPortfolioAssociationObservation struct {
}

type PrincipalPortfolioAssociationParameters struct {

	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// +kubebuilder:validation:Required
	PortfolioID *string `json:"portfolioId" tf:"portfolio_id,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalArn *string `json:"principalArn" tf:"principal_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PrincipalPortfolioAssociationSpec defines the desired state of PrincipalPortfolioAssociation
type PrincipalPortfolioAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrincipalPortfolioAssociationParameters `json:"forProvider"`
}

// PrincipalPortfolioAssociationStatus defines the observed state of PrincipalPortfolioAssociation.
type PrincipalPortfolioAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrincipalPortfolioAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrincipalPortfolioAssociation is the Schema for the PrincipalPortfolioAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type PrincipalPortfolioAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrincipalPortfolioAssociationSpec   `json:"spec"`
	Status            PrincipalPortfolioAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrincipalPortfolioAssociationList contains a list of PrincipalPortfolioAssociations
type PrincipalPortfolioAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrincipalPortfolioAssociation `json:"items"`
}

// Repository type metadata.
var (
	PrincipalPortfolioAssociationKind             = "PrincipalPortfolioAssociation"
	PrincipalPortfolioAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: PrincipalPortfolioAssociationKind}.String()
	PrincipalPortfolioAssociationKindAPIVersion   = PrincipalPortfolioAssociationKind + "." + GroupVersion.String()
	PrincipalPortfolioAssociationGroupVersionKind = GroupVersion.WithKind(PrincipalPortfolioAssociationKind)
)

func init() {
	SchemeBuilder.Register(&PrincipalPortfolioAssociation{}, &PrincipalPortfolioAssociationList{})
}