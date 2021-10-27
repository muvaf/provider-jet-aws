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

type ExcludesObservation struct {
}

type ExcludesParameters struct {

	// +kubebuilder:validation:Optional
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	Atime *string `json:"atime,omitempty" tf:"atime,omitempty"`

	// +kubebuilder:validation:Optional
	BytesPerSecond *int64 `json:"bytesPerSecond,omitempty" tf:"bytes_per_second,omitempty"`

	// +kubebuilder:validation:Optional
	GID *string `json:"gid,omitempty" tf:"gid,omitempty"`

	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// +kubebuilder:validation:Optional
	Mtime *string `json:"mtime,omitempty" tf:"mtime,omitempty"`

	// +kubebuilder:validation:Optional
	OverwriteMode *string `json:"overwriteMode,omitempty" tf:"overwrite_mode,omitempty"`

	// +kubebuilder:validation:Optional
	PosixPermissions *string `json:"posixPermissions,omitempty" tf:"posix_permissions,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveDeletedFiles *string `json:"preserveDeletedFiles,omitempty" tf:"preserve_deleted_files,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveDevices *string `json:"preserveDevices,omitempty" tf:"preserve_devices,omitempty"`

	// +kubebuilder:validation:Optional
	TaskQueueing *string `json:"taskQueueing,omitempty" tf:"task_queueing,omitempty"`

	// +kubebuilder:validation:Optional
	TransferMode *string `json:"transferMode,omitempty" tf:"transfer_mode,omitempty"`

	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// +kubebuilder:validation:Optional
	VerifyMode *string `json:"verifyMode,omitempty" tf:"verify_mode,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	ScheduleExpression *string `json:"scheduleExpression" tf:"schedule_expression,omitempty"`
}

type TaskObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TaskParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn,omitempty"`

	// +kubebuilder:validation:Required
	DestinationLocationArn *string `json:"destinationLocationArn" tf:"destination_location_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Excludes []ExcludesParameters `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Required
	SourceLocationArn *string `json:"sourceLocationArn" tf:"source_location_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TaskSpec defines the desired state of Task
type TaskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TaskParameters `json:"forProvider"`
}

// TaskStatus defines the observed state of Task.
type TaskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Task is the Schema for the Tasks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Task struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TaskSpec   `json:"spec"`
	Status            TaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TaskList contains a list of Tasks
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Task `json:"items"`
}

// Repository type metadata.
var (
	TaskKind             = "Task"
	TaskGroupKind        = schema.GroupKind{Group: Group, Kind: TaskKind}.String()
	TaskKindAPIVersion   = TaskKind + "." + GroupVersion.String()
	TaskGroupVersionKind = GroupVersion.WithKind(TaskKind)
)

func init() {
	SchemeBuilder.Register(&Task{}, &TaskList{})
}