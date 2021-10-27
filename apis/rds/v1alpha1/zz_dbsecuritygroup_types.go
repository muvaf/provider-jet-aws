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

type DbSecurityGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DbSecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Ingress []IngressParameters `json:"ingress" tf:"ingress,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IngressObservation struct {
}

type IngressParameters struct {

	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupOwnerID *string `json:"securityGroupOwnerId,omitempty" tf:"security_group_owner_id,omitempty"`
}

// DbSecurityGroupSpec defines the desired state of DbSecurityGroup
type DbSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DbSecurityGroupParameters `json:"forProvider"`
}

// DbSecurityGroupStatus defines the observed state of DbSecurityGroup.
type DbSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DbSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbSecurityGroup is the Schema for the DbSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSecurityGroupSpec   `json:"spec"`
	Status            DbSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSecurityGroupList contains a list of DbSecurityGroups
type DbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DbSecurityGroupKind             = "DbSecurityGroup"
	DbSecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: DbSecurityGroupKind}.String()
	DbSecurityGroupKindAPIVersion   = DbSecurityGroupKind + "." + GroupVersion.String()
	DbSecurityGroupGroupVersionKind = GroupVersion.WithKind(DbSecurityGroupKind)
)

func init() {
	SchemeBuilder.Register(&DbSecurityGroup{}, &DbSecurityGroupList{})
}