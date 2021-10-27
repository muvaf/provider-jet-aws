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

type AuthenticationOptionsObservation struct {
}

type AuthenticationOptionsParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectoryID *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id,omitempty"`

	// +kubebuilder:validation:Optional
	RootCertificateChainArn *string `json:"rootCertificateChainArn,omitempty" tf:"root_certificate_chain_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SamlProviderArn *string `json:"samlProviderArn,omitempty" tf:"saml_provider_arn,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClientVpnEndpointObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClientVpnEndpointParameters struct {

	// +kubebuilder:validation:Required
	AuthenticationOptions []AuthenticationOptionsParameters `json:"authenticationOptions" tf:"authentication_options,omitempty"`

	// +kubebuilder:validation:Required
	ClientCidrBlock *string `json:"clientCidrBlock" tf:"client_cidr_block,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionLogOptions []ConnectionLogOptionsParameters `json:"connectionLogOptions" tf:"connection_log_options,omitempty"`

	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServerCertificateArn *string `json:"serverCertificateArn" tf:"server_certificate_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SplitTunnel *bool `json:"splitTunnel,omitempty" tf:"split_tunnel,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol,omitempty"`
}

type ConnectionLogOptionsObservation struct {
}

type ConnectionLogOptionsParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogGroup *string `json:"cloudwatchLogGroup,omitempty" tf:"cloudwatch_log_group,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchLogStream *string `json:"cloudwatchLogStream,omitempty" tf:"cloudwatch_log_stream,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

// ClientVpnEndpointSpec defines the desired state of ClientVpnEndpoint
type ClientVpnEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientVpnEndpointParameters `json:"forProvider"`
}

// ClientVpnEndpointStatus defines the observed state of ClientVpnEndpoint.
type ClientVpnEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientVpnEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClientVpnEndpoint is the Schema for the ClientVpnEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ClientVpnEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientVpnEndpointSpec   `json:"spec"`
	Status            ClientVpnEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientVpnEndpointList contains a list of ClientVpnEndpoints
type ClientVpnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientVpnEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ClientVpnEndpointKind             = "ClientVpnEndpoint"
	ClientVpnEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: ClientVpnEndpointKind}.String()
	ClientVpnEndpointKindAPIVersion   = ClientVpnEndpointKind + "." + GroupVersion.String()
	ClientVpnEndpointGroupVersionKind = GroupVersion.WithKind(ClientVpnEndpointKind)
)

func init() {
	SchemeBuilder.Register(&ClientVpnEndpoint{}, &ClientVpnEndpointList{})
}