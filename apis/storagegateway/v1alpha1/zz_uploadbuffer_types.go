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

type UploadBufferObservation struct {
}

type UploadBufferParameters struct {

	// +kubebuilder:validation:Optional
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskPath *string `json:"diskPath,omitempty" tf:"disk_path,omitempty"`

	// +kubebuilder:validation:Required
	GatewayArn *string `json:"gatewayArn" tf:"gateway_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// UploadBufferSpec defines the desired state of UploadBuffer
type UploadBufferSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UploadBufferParameters `json:"forProvider"`
}

// UploadBufferStatus defines the observed state of UploadBuffer.
type UploadBufferStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UploadBufferObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UploadBuffer is the Schema for the UploadBuffers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type UploadBuffer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UploadBufferSpec   `json:"spec"`
	Status            UploadBufferStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UploadBufferList contains a list of UploadBuffers
type UploadBufferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UploadBuffer `json:"items"`
}

// Repository type metadata.
var (
	UploadBufferKind             = "UploadBuffer"
	UploadBufferGroupKind        = schema.GroupKind{Group: Group, Kind: UploadBufferKind}.String()
	UploadBufferKindAPIVersion   = UploadBufferKind + "." + GroupVersion.String()
	UploadBufferGroupVersionKind = GroupVersion.WithKind(UploadBufferKind)
)

func init() {
	SchemeBuilder.Register(&UploadBuffer{}, &UploadBufferList{})
}