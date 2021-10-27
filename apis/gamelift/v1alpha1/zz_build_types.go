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

type BuildObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type BuildParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OperatingSystem *string `json:"operatingSystem" tf:"operating_system,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StorageLocation []StorageLocationParameters `json:"storageLocation" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StorageLocationObservation struct {
}

type StorageLocationParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

// BuildSpec defines the desired state of Build
type BuildSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BuildParameters `json:"forProvider"`
}

// BuildStatus defines the observed state of Build.
type BuildStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BuildObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Build is the Schema for the Builds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BuildSpec   `json:"spec"`
	Status            BuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BuildList contains a list of Builds
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}

// Repository type metadata.
var (
	BuildKind             = "Build"
	BuildGroupKind        = schema.GroupKind{Group: Group, Kind: BuildKind}.String()
	BuildKindAPIVersion   = BuildKind + "." + GroupVersion.String()
	BuildGroupVersionKind = GroupVersion.WithKind(BuildKind)
)

func init() {
	SchemeBuilder.Register(&Build{}, &BuildList{})
}