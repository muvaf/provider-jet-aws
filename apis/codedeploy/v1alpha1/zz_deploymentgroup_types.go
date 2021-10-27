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

type AlarmConfigurationObservation struct {
}

type AlarmConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IgnorePollAlarmFailure *bool `json:"ignorePollAlarmFailure,omitempty" tf:"ignore_poll_alarm_failure,omitempty"`
}

type AutoRollbackConfigurationObservation struct {
}

type AutoRollbackConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`
}

type BlueGreenDeploymentConfigObservation struct {
}

type BlueGreenDeploymentConfigParameters struct {

	// +kubebuilder:validation:Optional
	DeploymentReadyOption []DeploymentReadyOptionParameters `json:"deploymentReadyOption,omitempty" tf:"deployment_ready_option,omitempty"`

	// +kubebuilder:validation:Optional
	GreenFleetProvisioningOption []GreenFleetProvisioningOptionParameters `json:"greenFleetProvisioningOption,omitempty" tf:"green_fleet_provisioning_option,omitempty"`

	// +kubebuilder:validation:Optional
	TerminateBlueInstancesOnDeploymentSuccess []TerminateBlueInstancesOnDeploymentSuccessParameters `json:"terminateBlueInstancesOnDeploymentSuccess,omitempty" tf:"terminate_blue_instances_on_deployment_success,omitempty"`
}

type DeploymentGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	DeploymentGroupID *string `json:"deploymentGroupId,omitempty" tf:"deployment_group_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeploymentGroupParameters struct {

	// +kubebuilder:validation:Optional
	AlarmConfiguration []AlarmConfigurationParameters `json:"alarmConfiguration,omitempty" tf:"alarm_configuration,omitempty"`

	// +kubebuilder:validation:Required
	AppName *string `json:"appName" tf:"app_name,omitempty"`

	// +kubebuilder:validation:Optional
	AutoRollbackConfiguration []AutoRollbackConfigurationParameters `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscalingGroups []*string `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`

	// +kubebuilder:validation:Optional
	BlueGreenDeploymentConfig []BlueGreenDeploymentConfigParameters `json:"blueGreenDeploymentConfig,omitempty" tf:"blue_green_deployment_config,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentConfigName *string `json:"deploymentConfigName,omitempty" tf:"deployment_config_name,omitempty"`

	// +kubebuilder:validation:Required
	DeploymentGroupName *string `json:"deploymentGroupName" tf:"deployment_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentStyle []DeploymentStyleParameters `json:"deploymentStyle,omitempty" tf:"deployment_style,omitempty"`

	// +kubebuilder:validation:Optional
	Ec2TagFilter []Ec2TagFilterParameters `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter,omitempty"`

	// +kubebuilder:validation:Optional
	Ec2TagSet []Ec2TagSetParameters `json:"ec2TagSet,omitempty" tf:"ec2_tag_set,omitempty"`

	// +kubebuilder:validation:Optional
	EcsService []EcsServiceParameters `json:"ecsService,omitempty" tf:"ecs_service,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerInfo []LoadBalancerInfoParameters `json:"loadBalancerInfo,omitempty" tf:"load_balancer_info,omitempty"`

	// +kubebuilder:validation:Optional
	OnPremisesInstanceTagFilter []OnPremisesInstanceTagFilterParameters `json:"onPremisesInstanceTagFilter,omitempty" tf:"on_premises_instance_tag_filter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServiceRoleArn *string `json:"serviceRoleArn" tf:"service_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TriggerConfiguration []TriggerConfigurationParameters `json:"triggerConfiguration,omitempty" tf:"trigger_configuration,omitempty"`
}

type DeploymentReadyOptionObservation struct {
}

type DeploymentReadyOptionParameters struct {

	// +kubebuilder:validation:Optional
	ActionOnTimeout *string `json:"actionOnTimeout,omitempty" tf:"action_on_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	WaitTimeInMinutes *int64 `json:"waitTimeInMinutes,omitempty" tf:"wait_time_in_minutes,omitempty"`
}

type DeploymentStyleObservation struct {
}

type DeploymentStyleParameters struct {

	// +kubebuilder:validation:Optional
	DeploymentOption *string `json:"deploymentOption,omitempty" tf:"deployment_option,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`
}

type Ec2TagFilterObservation struct {
}

type Ec2TagFilterParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type Ec2TagSetEc2TagFilterObservation struct {
}

type Ec2TagSetEc2TagFilterParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type Ec2TagSetObservation struct {
}

type Ec2TagSetParameters struct {

	// +kubebuilder:validation:Optional
	Ec2TagFilter []Ec2TagSetEc2TagFilterParameters `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter,omitempty"`
}

type EcsServiceObservation struct {
}

type EcsServiceParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

type ElbInfoObservation struct {
}

type ElbInfoParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GreenFleetProvisioningOptionObservation struct {
}

type GreenFleetProvisioningOptionParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`
}

type LoadBalancerInfoObservation struct {
}

type LoadBalancerInfoParameters struct {

	// +kubebuilder:validation:Optional
	ElbInfo []ElbInfoParameters `json:"elbInfo,omitempty" tf:"elb_info,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupInfo []TargetGroupInfoParameters `json:"targetGroupInfo,omitempty" tf:"target_group_info,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupPairInfo []TargetGroupPairInfoParameters `json:"targetGroupPairInfo,omitempty" tf:"target_group_pair_info,omitempty"`
}

type OnPremisesInstanceTagFilterObservation struct {
}

type OnPremisesInstanceTagFilterParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProdTrafficRouteObservation struct {
}

type ProdTrafficRouteParameters struct {

	// +kubebuilder:validation:Required
	ListenerArns []*string `json:"listenerArns" tf:"listener_arns,omitempty"`
}

type TargetGroupInfoObservation struct {
}

type TargetGroupInfoParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TargetGroupObservation struct {
}

type TargetGroupPairInfoObservation struct {
}

type TargetGroupPairInfoParameters struct {

	// +kubebuilder:validation:Required
	ProdTrafficRoute []ProdTrafficRouteParameters `json:"prodTrafficRoute" tf:"prod_traffic_route,omitempty"`

	// +kubebuilder:validation:Required
	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group,omitempty"`

	// +kubebuilder:validation:Optional
	TestTrafficRoute []TestTrafficRouteParameters `json:"testTrafficRoute,omitempty" tf:"test_traffic_route,omitempty"`
}

type TargetGroupParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type TerminateBlueInstancesOnDeploymentSuccessObservation struct {
}

type TerminateBlueInstancesOnDeploymentSuccessParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationWaitTimeInMinutes *int64 `json:"terminationWaitTimeInMinutes,omitempty" tf:"termination_wait_time_in_minutes,omitempty"`
}

type TestTrafficRouteObservation struct {
}

type TestTrafficRouteParameters struct {

	// +kubebuilder:validation:Required
	ListenerArns []*string `json:"listenerArns" tf:"listener_arns,omitempty"`
}

type TriggerConfigurationObservation struct {
}

type TriggerConfigurationParameters struct {

	// +kubebuilder:validation:Required
	TriggerEvents []*string `json:"triggerEvents" tf:"trigger_events,omitempty"`

	// +kubebuilder:validation:Required
	TriggerName *string `json:"triggerName" tf:"trigger_name,omitempty"`

	// +kubebuilder:validation:Required
	TriggerTargetArn *string `json:"triggerTargetArn" tf:"trigger_target_arn,omitempty"`
}

// DeploymentGroupSpec defines the desired state of DeploymentGroup
type DeploymentGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentGroupParameters `json:"forProvider"`
}

// DeploymentGroupStatus defines the observed state of DeploymentGroup.
type DeploymentGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentGroup is the Schema for the DeploymentGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentGroupSpec   `json:"spec"`
	Status            DeploymentGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentGroupList contains a list of DeploymentGroups
type DeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentGroup `json:"items"`
}

// Repository type metadata.
var (
	DeploymentGroupKind             = "DeploymentGroup"
	DeploymentGroupGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentGroupKind}.String()
	DeploymentGroupKindAPIVersion   = DeploymentGroupKind + "." + GroupVersion.String()
	DeploymentGroupGroupVersionKind = GroupVersion.WithKind(DeploymentGroupKind)
)

func init() {
	SchemeBuilder.Register(&DeploymentGroup{}, &DeploymentGroupList{})
}