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

type LogMetricFilterObservation struct {
}

type LogMetricFilterParameters struct {

	// +kubebuilder:validation:Required
	LogGroupName *string `json:"logGroupName" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Required
	MetricTransformation []MetricTransformationParameters `json:"metricTransformation" tf:"metric_transformation,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Pattern *string `json:"pattern" tf:"pattern,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type MetricTransformationObservation struct {
}

type MetricTransformationParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// LogMetricFilterSpec defines the desired state of LogMetricFilter
type LogMetricFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogMetricFilterParameters `json:"forProvider"`
}

// LogMetricFilterStatus defines the observed state of LogMetricFilter.
type LogMetricFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogMetricFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogMetricFilter is the Schema for the LogMetricFilters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LogMetricFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogMetricFilterSpec   `json:"spec"`
	Status            LogMetricFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogMetricFilterList contains a list of LogMetricFilters
type LogMetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogMetricFilter `json:"items"`
}

// Repository type metadata.
var (
	LogMetricFilterKind             = "LogMetricFilter"
	LogMetricFilterGroupKind        = schema.GroupKind{Group: Group, Kind: LogMetricFilterKind}.String()
	LogMetricFilterKindAPIVersion   = LogMetricFilterKind + "." + GroupVersion.String()
	LogMetricFilterGroupVersionKind = GroupVersion.WithKind(LogMetricFilterKind)
)

func init() {
	SchemeBuilder.Register(&LogMetricFilter{}, &LogMetricFilterList{})
}