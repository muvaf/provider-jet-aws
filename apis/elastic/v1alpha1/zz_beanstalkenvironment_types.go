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

type AllSettingsObservation struct {
}

type AllSettingsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BeanstalkEnvironmentObservation struct {
	AllSettings []AllSettingsObservation `json:"allSettings,omitempty" tf:"all_settings,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AutoscalingGroups []*string `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`

	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	EndpointURL *string `json:"endpointUrl,omitempty" tf:"endpoint_url,omitempty"`

	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	LaunchConfigurations []*string `json:"launchConfigurations,omitempty" tf:"launch_configurations,omitempty"`

	LoadBalancers []*string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`

	Queues []*string `json:"queues,omitempty" tf:"queues,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Triggers []*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type BeanstalkEnvironmentParameters struct {

	// +kubebuilder:validation:Required
	Application *string `json:"application" tf:"application,omitempty"`

	// +kubebuilder:validation:Optional
	CnamePrefix *string `json:"cnamePrefix,omitempty" tf:"cname_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformArn *string `json:"platformArn,omitempty" tf:"platform_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PollInterval *string `json:"pollInterval,omitempty" tf:"poll_interval,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Setting []BeanstalkEnvironmentSettingParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// +kubebuilder:validation:Optional
	SolutionStackName *string `json:"solutionStackName,omitempty" tf:"solution_stack_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// +kubebuilder:validation:Optional
	VersionLabel *string `json:"versionLabel,omitempty" tf:"version_label,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForReadyTimeout *string `json:"waitForReadyTimeout,omitempty" tf:"wait_for_ready_timeout,omitempty"`
}

type BeanstalkEnvironmentSettingObservation struct {
}

type BeanstalkEnvironmentSettingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// BeanstalkEnvironmentSpec defines the desired state of BeanstalkEnvironment
type BeanstalkEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BeanstalkEnvironmentParameters `json:"forProvider"`
}

// BeanstalkEnvironmentStatus defines the observed state of BeanstalkEnvironment.
type BeanstalkEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BeanstalkEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BeanstalkEnvironment is the Schema for the BeanstalkEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BeanstalkEnvironmentSpec   `json:"spec"`
	Status            BeanstalkEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BeanstalkEnvironmentList contains a list of BeanstalkEnvironments
type BeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeanstalkEnvironment `json:"items"`
}

// Repository type metadata.
var (
	BeanstalkEnvironmentKind             = "BeanstalkEnvironment"
	BeanstalkEnvironmentGroupKind        = schema.GroupKind{Group: Group, Kind: BeanstalkEnvironmentKind}.String()
	BeanstalkEnvironmentKindAPIVersion   = BeanstalkEnvironmentKind + "." + GroupVersion.String()
	BeanstalkEnvironmentGroupVersionKind = GroupVersion.WithKind(BeanstalkEnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&BeanstalkEnvironment{}, &BeanstalkEnvironmentList{})
}