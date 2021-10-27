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

type GatewayNetworkInterfaceObservation struct {
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`
}

type GatewayNetworkInterfaceParameters struct {
}

type GatewayObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Ec2InstanceID *string `json:"ec2InstanceId,omitempty" tf:"ec2_instance_id,omitempty"`

	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	GatewayNetworkInterface []GatewayNetworkInterfaceObservation `json:"gatewayNetworkInterface,omitempty" tf:"gateway_network_interface,omitempty"`

	HostEnvironment *string `json:"hostEnvironment,omitempty" tf:"host_environment,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GatewayParameters struct {

	// +kubebuilder:validation:Optional
	ActivationKey *string `json:"activationKey,omitempty" tf:"activation_key,omitempty"`

	// +kubebuilder:validation:Optional
	AverageDownloadRateLimitInBitsPerSec *int64 `json:"averageDownloadRateLimitInBitsPerSec,omitempty" tf:"average_download_rate_limit_in_bits_per_sec,omitempty"`

	// +kubebuilder:validation:Optional
	AverageUploadRateLimitInBitsPerSec *int64 `json:"averageUploadRateLimitInBitsPerSec,omitempty" tf:"average_upload_rate_limit_in_bits_per_sec,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayIPAddress *string `json:"gatewayIpAddress,omitempty" tf:"gateway_ip_address,omitempty"`

	// +kubebuilder:validation:Required
	GatewayName *string `json:"gatewayName" tf:"gateway_name,omitempty"`

	// +kubebuilder:validation:Required
	GatewayTimezone *string `json:"gatewayTimezone" tf:"gateway_timezone,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayVpcEndpoint *string `json:"gatewayVpcEndpoint,omitempty" tf:"gateway_vpc_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	MediumChangerType *string `json:"mediumChangerType,omitempty" tf:"medium_changer_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SmbActiveDirectorySettings []SmbActiveDirectorySettingsParameters `json:"smbActiveDirectorySettings,omitempty" tf:"smb_active_directory_settings,omitempty"`

	// +kubebuilder:validation:Optional
	SmbFileShareVisibility *bool `json:"smbFileShareVisibility,omitempty" tf:"smb_file_share_visibility,omitempty"`

	// +kubebuilder:validation:Optional
	SmbGuestPasswordSecretRef v1.SecretKeySelector `json:"smbGuestPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SmbSecurityStrategy *string `json:"smbSecurityStrategy,omitempty" tf:"smb_security_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TapeDriveType *string `json:"tapeDriveType,omitempty" tf:"tape_drive_type,omitempty"`
}

type SmbActiveDirectorySettingsObservation struct {
	ActiveDirectoryStatus *string `json:"activeDirectoryStatus,omitempty" tf:"active_directory_status,omitempty"`
}

type SmbActiveDirectorySettingsParameters struct {

	// +kubebuilder:validation:Optional
	DomainControllers []*string `json:"domainControllers,omitempty" tf:"domain_controllers,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	GatewayKind             = "Gateway"
	GatewayGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayKind}.String()
	GatewayKindAPIVersion   = GatewayKind + "." + GroupVersion.String()
	GatewayGroupVersionKind = GroupVersion.WithKind(GatewayKind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}