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

type BucketPublicAccessBlockObservation struct {
}

type BucketPublicAccessBlockParameters struct {

	// +kubebuilder:validation:Optional
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

// BucketPublicAccessBlockSpec defines the desired state of BucketPublicAccessBlock
type BucketPublicAccessBlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketPublicAccessBlockParameters `json:"forProvider"`
}

// BucketPublicAccessBlockStatus defines the observed state of BucketPublicAccessBlock.
type BucketPublicAccessBlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketPublicAccessBlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketPublicAccessBlock is the Schema for the BucketPublicAccessBlocks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BucketPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketPublicAccessBlockSpec   `json:"spec"`
	Status            BucketPublicAccessBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketPublicAccessBlockList contains a list of BucketPublicAccessBlocks
type BucketPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketPublicAccessBlock `json:"items"`
}

// Repository type metadata.
var (
	BucketPublicAccessBlockKind             = "BucketPublicAccessBlock"
	BucketPublicAccessBlockGroupKind        = schema.GroupKind{Group: Group, Kind: BucketPublicAccessBlockKind}.String()
	BucketPublicAccessBlockKindAPIVersion   = BucketPublicAccessBlockKind + "." + GroupVersion.String()
	BucketPublicAccessBlockGroupVersionKind = GroupVersion.WithKind(BucketPublicAccessBlockKind)
)

func init() {
	SchemeBuilder.Register(&BucketPublicAccessBlock{}, &BucketPublicAccessBlockList{})
}