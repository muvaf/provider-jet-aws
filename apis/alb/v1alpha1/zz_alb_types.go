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

type AccessLogsObservation struct {
}

type AccessLogsParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type AlbObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AlbParameters struct {

	// +kubebuilder:validation:Optional
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// +kubebuilder:validation:Optional
	DropInvalidHeaderFields *bool `json:"dropInvalidHeaderFields,omitempty" tf:"drop_invalid_header_fields,omitempty"`

	// +kubebuilder:validation:Optional
	EnableCrossZoneLoadBalancing *bool `json:"enableCrossZoneLoadBalancing,omitempty" tf:"enable_cross_zone_load_balancing,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDeletionProtection *bool `json:"enableDeletionProtection,omitempty" tf:"enable_deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeout *int64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetMapping []SubnetMappingParameters `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SubnetMappingObservation struct {
	OutpostID *string `json:"outpostId,omitempty" tf:"outpost_id,omitempty"`
}

type SubnetMappingParameters struct {

	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPv4Address *string `json:"privateIpv4Address,omitempty" tf:"private_ipv4_address,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// AlbSpec defines the desired state of Alb
type AlbSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlbParameters `json:"forProvider"`
}

// AlbStatus defines the observed state of Alb.
type AlbStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlbObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alb is the Schema for the Albs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Alb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbSpec   `json:"spec"`
	Status            AlbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbList contains a list of Albs
type AlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alb `json:"items"`
}

// Repository type metadata.
var (
	AlbKind             = "Alb"
	AlbGroupKind        = schema.GroupKind{Group: Group, Kind: AlbKind}.String()
	AlbKindAPIVersion   = AlbKind + "." + GroupVersion.String()
	AlbGroupVersionKind = GroupVersion.WithKind(AlbKind)
)

func init() {
	SchemeBuilder.Register(&Alb{}, &AlbList{})
}