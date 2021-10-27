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

type AdvancedSecurityOptionsObservation struct {
}

type AdvancedSecurityOptionsParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty" tf:"internal_user_database_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MasterUserOptions []MasterUserOptionsParameters `json:"masterUserOptions,omitempty" tf:"master_user_options,omitempty"`
}

type ClusterConfigObservation struct {
}

type ClusterConfigParameters struct {

	// +kubebuilder:validation:Optional
	DedicatedMasterCount *int64 `json:"dedicatedMasterCount,omitempty" tf:"dedicated_master_count,omitempty"`

	// +kubebuilder:validation:Optional
	DedicatedMasterEnabled *bool `json:"dedicatedMasterEnabled,omitempty" tf:"dedicated_master_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DedicatedMasterType *string `json:"dedicatedMasterType,omitempty" tf:"dedicated_master_type,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	WarmCount *int64 `json:"warmCount,omitempty" tf:"warm_count,omitempty"`

	// +kubebuilder:validation:Optional
	WarmEnabled *bool `json:"warmEnabled,omitempty" tf:"warm_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	WarmType *string `json:"warmType,omitempty" tf:"warm_type,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneAwarenessConfig []ZoneAwarenessConfigParameters `json:"zoneAwarenessConfig,omitempty" tf:"zone_awareness_config,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneAwarenessEnabled *bool `json:"zoneAwarenessEnabled,omitempty" tf:"zone_awareness_enabled,omitempty"`
}

type CognitoOptionsObservation struct {
}

type CognitoOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	IdentityPoolID *string `json:"identityPoolId" tf:"identity_pool_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Required
	UserPoolID *string `json:"userPoolId" tf:"user_pool_id,omitempty"`
}

type DomainEndpointOptionsObservation struct {
}

type DomainEndpointOptionsParameters struct {

	// +kubebuilder:validation:Optional
	CustomEndpoint *string `json:"customEndpoint,omitempty" tf:"custom_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn,omitempty" tf:"custom_endpoint_certificate_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CustomEndpointEnabled *bool `json:"customEndpointEnabled,omitempty" tf:"custom_endpoint_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// +kubebuilder:validation:Optional
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type DomainObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	KibanaEndpoint *string `json:"kibanaEndpoint,omitempty" tf:"kibana_endpoint,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DomainParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicies *string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`

	// +kubebuilder:validation:Optional
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty" tf:"advanced_options,omitempty"`

	// +kubebuilder:validation:Optional
	AdvancedSecurityOptions []AdvancedSecurityOptionsParameters `json:"advancedSecurityOptions,omitempty" tf:"advanced_security_options,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterConfig []ClusterConfigParameters `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`

	// +kubebuilder:validation:Optional
	CognitoOptions []CognitoOptionsParameters `json:"cognitoOptions,omitempty" tf:"cognito_options,omitempty"`

	// +kubebuilder:validation:Optional
	DomainEndpointOptions []DomainEndpointOptionsParameters `json:"domainEndpointOptions,omitempty" tf:"domain_endpoint_options,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	EbsOptions []EbsOptionsParameters `json:"ebsOptions,omitempty" tf:"ebs_options,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty" tf:"elasticsearch_version,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptAtRest []EncryptAtRestParameters `json:"encryptAtRest,omitempty" tf:"encrypt_at_rest,omitempty"`

	// +kubebuilder:validation:Optional
	LogPublishingOptions []LogPublishingOptionsParameters `json:"logPublishingOptions,omitempty" tf:"log_publishing_options,omitempty"`

	// +kubebuilder:validation:Optional
	NodeToNodeEncryption []NodeToNodeEncryptionParameters `json:"nodeToNodeEncryption,omitempty" tf:"node_to_node_encryption,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SnapshotOptions []SnapshotOptionsParameters `json:"snapshotOptions,omitempty" tf:"snapshot_options,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VpcOptions []VpcOptionsParameters `json:"vpcOptions,omitempty" tf:"vpc_options,omitempty"`
}

type EbsOptionsObservation struct {
}

type EbsOptionsParameters struct {

	// +kubebuilder:validation:Required
	EbsEnabled *bool `json:"ebsEnabled" tf:"ebs_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EncryptAtRestObservation struct {
}

type EncryptAtRestParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type LogPublishingOptionsObservation struct {
}

type LogPublishingOptionsParameters struct {

	// +kubebuilder:validation:Required
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn" tf:"cloudwatch_log_group_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

type MasterUserOptionsObservation struct {
}

type MasterUserOptionsParameters struct {

	// +kubebuilder:validation:Optional
	MasterUserArn *string `json:"masterUserArn,omitempty" tf:"master_user_arn,omitempty"`

	// +kubebuilder:validation:Optional
	MasterUserName *string `json:"masterUserName,omitempty" tf:"master_user_name,omitempty"`

	// +kubebuilder:validation:Optional
	MasterUserPasswordSecretRef v1.SecretKeySelector `json:"masterUserPasswordSecretRef,omitempty" tf:"-"`
}

type NodeToNodeEncryptionObservation struct {
}

type NodeToNodeEncryptionParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SnapshotOptionsObservation struct {
}

type SnapshotOptionsParameters struct {

	// +kubebuilder:validation:Required
	AutomatedSnapshotStartHour *int64 `json:"automatedSnapshotStartHour" tf:"automated_snapshot_start_hour,omitempty"`
}

type VpcOptionsObservation struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VpcOptionsParameters struct {

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type ZoneAwarenessConfigObservation struct {
}

type ZoneAwarenessConfigParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZoneCount *int64 `json:"availabilityZoneCount,omitempty" tf:"availability_zone_count,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	DomainKind             = "Domain"
	DomainGroupKind        = schema.GroupKind{Group: Group, Kind: DomainKind}.String()
	DomainKindAPIVersion   = DomainKind + "." + GroupVersion.String()
	DomainGroupVersionKind = GroupVersion.WithKind(DomainKind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}