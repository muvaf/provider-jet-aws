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

type GatewayMethodResponseObservation struct {
}

type GatewayMethodResponseParameters struct {

	// +kubebuilder:validation:Required
	HTTPMethod *string `json:"httpMethod" tf:"http_method,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`

	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

// GatewayMethodResponseSpec defines the desired state of GatewayMethodResponse
type GatewayMethodResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayMethodResponseParameters `json:"forProvider"`
}

// GatewayMethodResponseStatus defines the observed state of GatewayMethodResponse.
type GatewayMethodResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayMethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayMethodResponse is the Schema for the GatewayMethodResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayMethodResponseSpec   `json:"spec"`
	Status            GatewayMethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayMethodResponseList contains a list of GatewayMethodResponses
type GatewayMethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayMethodResponse `json:"items"`
}

// Repository type metadata.
var (
	GatewayMethodResponseKind             = "GatewayMethodResponse"
	GatewayMethodResponseGroupKind        = schema.GroupKind{Group: Group, Kind: GatewayMethodResponseKind}.String()
	GatewayMethodResponseKindAPIVersion   = GatewayMethodResponseKind + "." + GroupVersion.String()
	GatewayMethodResponseGroupVersionKind = GroupVersion.WithKind(GatewayMethodResponseKind)
)

func init() {
	SchemeBuilder.Register(&GatewayMethodResponse{}, &GatewayMethodResponseList{})
}