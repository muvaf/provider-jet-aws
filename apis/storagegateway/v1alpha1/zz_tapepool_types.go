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

type TapePoolObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TapePoolParameters struct {

	// +kubebuilder:validation:Required
	PoolName *string `json:"poolName" tf:"pool_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetentionLockTimeInDays *int64 `json:"retentionLockTimeInDays,omitempty" tf:"retention_lock_time_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionLockType *string `json:"retentionLockType,omitempty" tf:"retention_lock_type,omitempty"`

	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TapePoolSpec defines the desired state of TapePool
type TapePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TapePoolParameters `json:"forProvider"`
}

// TapePoolStatus defines the observed state of TapePool.
type TapePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TapePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TapePool is the Schema for the TapePools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TapePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TapePoolSpec   `json:"spec"`
	Status            TapePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TapePoolList contains a list of TapePools
type TapePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TapePool `json:"items"`
}

// Repository type metadata.
var (
	TapePoolKind             = "TapePool"
	TapePoolGroupKind        = schema.GroupKind{Group: Group, Kind: TapePoolKind}.String()
	TapePoolKindAPIVersion   = TapePoolKind + "." + GroupVersion.String()
	TapePoolGroupVersionKind = GroupVersion.WithKind(TapePoolKind)
)

func init() {
	SchemeBuilder.Register(&TapePool{}, &TapePoolList{})
}