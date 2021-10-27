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

type LocationS3Observation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationS3Parameters struct {

	// +kubebuilder:validation:Optional
	AgentArns []*string `json:"agentArns,omitempty" tf:"agent_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	S3BucketArn *string `json:"s3BucketArn" tf:"s3_bucket_arn,omitempty"`

	// +kubebuilder:validation:Required
	S3Config []S3ConfigParameters `json:"s3Config" tf:"s3_config,omitempty"`

	// +kubebuilder:validation:Optional
	S3StorageClass *string `json:"s3StorageClass,omitempty" tf:"s3_storage_class,omitempty"`

	// +kubebuilder:validation:Required
	Subdirectory *string `json:"subdirectory" tf:"subdirectory,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3ConfigObservation struct {
}

type S3ConfigParameters struct {

	// +kubebuilder:validation:Required
	BucketAccessRoleArn *string `json:"bucketAccessRoleArn" tf:"bucket_access_role_arn,omitempty"`
}

// LocationS3Spec defines the desired state of LocationS3
type LocationS3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationS3Parameters `json:"forProvider"`
}

// LocationS3Status defines the observed state of LocationS3.
type LocationS3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationS3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocationS3 is the Schema for the LocationS3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LocationS3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationS3Spec   `json:"spec"`
	Status            LocationS3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationS3List contains a list of LocationS3s
type LocationS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationS3 `json:"items"`
}

// Repository type metadata.
var (
	LocationS3Kind             = "LocationS3"
	LocationS3GroupKind        = schema.GroupKind{Group: Group, Kind: LocationS3Kind}.String()
	LocationS3KindAPIVersion   = LocationS3Kind + "." + GroupVersion.String()
	LocationS3GroupVersionKind = GroupVersion.WithKind(LocationS3Kind)
)

func init() {
	SchemeBuilder.Register(&LocationS3{}, &LocationS3List{})
}