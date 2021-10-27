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

type DbProxyEndpointObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DbProxyEndpointParameters struct {

	// +kubebuilder:validation:Required
	DBProxyEndpointName *string `json:"dbProxyEndpointName" tf:"db_proxy_endpoint_name,omitempty"`

	// +kubebuilder:validation:Required
	DBProxyName *string `json:"dbProxyName" tf:"db_proxy_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetRole *string `json:"targetRole,omitempty" tf:"target_role,omitempty"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	VpcSubnetIds []*string `json:"vpcSubnetIds" tf:"vpc_subnet_ids,omitempty"`
}

// DbProxyEndpointSpec defines the desired state of DbProxyEndpoint
type DbProxyEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DbProxyEndpointParameters `json:"forProvider"`
}

// DbProxyEndpointStatus defines the observed state of DbProxyEndpoint.
type DbProxyEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DbProxyEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyEndpoint is the Schema for the DbProxyEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbProxyEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbProxyEndpointSpec   `json:"spec"`
	Status            DbProxyEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyEndpointList contains a list of DbProxyEndpoints
type DbProxyEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxyEndpoint `json:"items"`
}

// Repository type metadata.
var (
	DbProxyEndpointKind             = "DbProxyEndpoint"
	DbProxyEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: DbProxyEndpointKind}.String()
	DbProxyEndpointKindAPIVersion   = DbProxyEndpointKind + "." + GroupVersion.String()
	DbProxyEndpointGroupVersionKind = GroupVersion.WithKind(DbProxyEndpointKind)
)

func init() {
	SchemeBuilder.Register(&DbProxyEndpoint{}, &DbProxyEndpointList{})
}