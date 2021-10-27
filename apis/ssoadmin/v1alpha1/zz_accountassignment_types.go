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

type AccountAssignmentObservation struct {
}

type AccountAssignmentParameters struct {

	// +kubebuilder:validation:Required
	InstanceArn *string `json:"instanceArn" tf:"instance_arn,omitempty"`

	// +kubebuilder:validation:Required
	PermissionSetArn *string `json:"permissionSetArn" tf:"permission_set_arn,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalType *string `json:"principalType" tf:"principal_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	TargetID *string `json:"targetId" tf:"target_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}

// AccountAssignmentSpec defines the desired state of AccountAssignment
type AccountAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountAssignmentParameters `json:"forProvider"`
}

// AccountAssignmentStatus defines the observed state of AccountAssignment.
type AccountAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAssignment is the Schema for the AccountAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AccountAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountAssignmentSpec   `json:"spec"`
	Status            AccountAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAssignmentList contains a list of AccountAssignments
type AccountAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountAssignment `json:"items"`
}

// Repository type metadata.
var (
	AccountAssignmentKind             = "AccountAssignment"
	AccountAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: AccountAssignmentKind}.String()
	AccountAssignmentKindAPIVersion   = AccountAssignmentKind + "." + GroupVersion.String()
	AccountAssignmentGroupVersionKind = GroupVersion.WithKind(AccountAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&AccountAssignment{}, &AccountAssignmentList{})
}