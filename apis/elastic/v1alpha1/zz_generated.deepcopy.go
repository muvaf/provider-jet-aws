// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllSettingsObservation) DeepCopyInto(out *AllSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllSettingsObservation.
func (in *AllSettingsObservation) DeepCopy() *AllSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(AllSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllSettingsParameters) DeepCopyInto(out *AllSettingsParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllSettingsParameters.
func (in *AllSettingsParameters) DeepCopy() *AllSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(AllSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppversionLifecycleObservation) DeepCopyInto(out *AppversionLifecycleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppversionLifecycleObservation.
func (in *AppversionLifecycleObservation) DeepCopy() *AppversionLifecycleObservation {
	if in == nil {
		return nil
	}
	out := new(AppversionLifecycleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppversionLifecycleParameters) DeepCopyInto(out *AppversionLifecycleParameters) {
	*out = *in
	if in.DeleteSourceFromS3 != nil {
		in, out := &in.DeleteSourceFromS3, &out.DeleteSourceFromS3
		*out = new(bool)
		**out = **in
	}
	if in.MaxAgeInDays != nil {
		in, out := &in.MaxAgeInDays, &out.MaxAgeInDays
		*out = new(int64)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int64)
		**out = **in
	}
	if in.ServiceRole != nil {
		in, out := &in.ServiceRole, &out.ServiceRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppversionLifecycleParameters.
func (in *AppversionLifecycleParameters) DeepCopy() *AppversionLifecycleParameters {
	if in == nil {
		return nil
	}
	out := new(AppversionLifecycleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplication) DeepCopyInto(out *BeanstalkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplication.
func (in *BeanstalkApplication) DeepCopy() *BeanstalkApplication {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationList) DeepCopyInto(out *BeanstalkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BeanstalkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationList.
func (in *BeanstalkApplicationList) DeepCopy() *BeanstalkApplicationList {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationObservation) DeepCopyInto(out *BeanstalkApplicationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationObservation.
func (in *BeanstalkApplicationObservation) DeepCopy() *BeanstalkApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationParameters) DeepCopyInto(out *BeanstalkApplicationParameters) {
	*out = *in
	if in.AppversionLifecycle != nil {
		in, out := &in.AppversionLifecycle, &out.AppversionLifecycle
		*out = make([]AppversionLifecycleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationParameters.
func (in *BeanstalkApplicationParameters) DeepCopy() *BeanstalkApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationSpec) DeepCopyInto(out *BeanstalkApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationSpec.
func (in *BeanstalkApplicationSpec) DeepCopy() *BeanstalkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationStatus) DeepCopyInto(out *BeanstalkApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationStatus.
func (in *BeanstalkApplicationStatus) DeepCopy() *BeanstalkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersion) DeepCopyInto(out *BeanstalkApplicationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersion.
func (in *BeanstalkApplicationVersion) DeepCopy() *BeanstalkApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkApplicationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersionList) DeepCopyInto(out *BeanstalkApplicationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BeanstalkApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersionList.
func (in *BeanstalkApplicationVersionList) DeepCopy() *BeanstalkApplicationVersionList {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkApplicationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersionObservation) DeepCopyInto(out *BeanstalkApplicationVersionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersionObservation.
func (in *BeanstalkApplicationVersionObservation) DeepCopy() *BeanstalkApplicationVersionObservation {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersionParameters) DeepCopyInto(out *BeanstalkApplicationVersionParameters) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersionParameters.
func (in *BeanstalkApplicationVersionParameters) DeepCopy() *BeanstalkApplicationVersionParameters {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersionSpec) DeepCopyInto(out *BeanstalkApplicationVersionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersionSpec.
func (in *BeanstalkApplicationVersionSpec) DeepCopy() *BeanstalkApplicationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkApplicationVersionStatus) DeepCopyInto(out *BeanstalkApplicationVersionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkApplicationVersionStatus.
func (in *BeanstalkApplicationVersionStatus) DeepCopy() *BeanstalkApplicationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(BeanstalkApplicationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplate) DeepCopyInto(out *BeanstalkConfigurationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplate.
func (in *BeanstalkConfigurationTemplate) DeepCopy() *BeanstalkConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkConfigurationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplateList) DeepCopyInto(out *BeanstalkConfigurationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BeanstalkConfigurationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplateList.
func (in *BeanstalkConfigurationTemplateList) DeepCopy() *BeanstalkConfigurationTemplateList {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkConfigurationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplateObservation) DeepCopyInto(out *BeanstalkConfigurationTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplateObservation.
func (in *BeanstalkConfigurationTemplateObservation) DeepCopy() *BeanstalkConfigurationTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplateParameters) DeepCopyInto(out *BeanstalkConfigurationTemplateParameters) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentID != nil {
		in, out := &in.EnvironmentID, &out.EnvironmentID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Setting != nil {
		in, out := &in.Setting, &out.Setting
		*out = make([]SettingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SolutionStackName != nil {
		in, out := &in.SolutionStackName, &out.SolutionStackName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplateParameters.
func (in *BeanstalkConfigurationTemplateParameters) DeepCopy() *BeanstalkConfigurationTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplateSpec) DeepCopyInto(out *BeanstalkConfigurationTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplateSpec.
func (in *BeanstalkConfigurationTemplateSpec) DeepCopy() *BeanstalkConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkConfigurationTemplateStatus) DeepCopyInto(out *BeanstalkConfigurationTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkConfigurationTemplateStatus.
func (in *BeanstalkConfigurationTemplateStatus) DeepCopy() *BeanstalkConfigurationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(BeanstalkConfigurationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironment) DeepCopyInto(out *BeanstalkEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironment.
func (in *BeanstalkEnvironment) DeepCopy() *BeanstalkEnvironment {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentList) DeepCopyInto(out *BeanstalkEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BeanstalkEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentList.
func (in *BeanstalkEnvironmentList) DeepCopy() *BeanstalkEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BeanstalkEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentObservation) DeepCopyInto(out *BeanstalkEnvironmentObservation) {
	*out = *in
	if in.AllSettings != nil {
		in, out := &in.AllSettings, &out.AllSettings
		*out = make([]AllSettingsObservation, len(*in))
		copy(*out, *in)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoscalingGroups != nil {
		in, out := &in.AutoscalingGroups, &out.AutoscalingGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LaunchConfigurations != nil {
		in, out := &in.LaunchConfigurations, &out.LaunchConfigurations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentObservation.
func (in *BeanstalkEnvironmentObservation) DeepCopy() *BeanstalkEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentParameters) DeepCopyInto(out *BeanstalkEnvironmentParameters) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(string)
		**out = **in
	}
	if in.CnamePrefix != nil {
		in, out := &in.CnamePrefix, &out.CnamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PlatformArn != nil {
		in, out := &in.PlatformArn, &out.PlatformArn
		*out = new(string)
		**out = **in
	}
	if in.PollInterval != nil {
		in, out := &in.PollInterval, &out.PollInterval
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Setting != nil {
		in, out := &in.Setting, &out.Setting
		*out = make([]BeanstalkEnvironmentSettingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SolutionStackName != nil {
		in, out := &in.SolutionStackName, &out.SolutionStackName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.VersionLabel != nil {
		in, out := &in.VersionLabel, &out.VersionLabel
		*out = new(string)
		**out = **in
	}
	if in.WaitForReadyTimeout != nil {
		in, out := &in.WaitForReadyTimeout, &out.WaitForReadyTimeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentParameters.
func (in *BeanstalkEnvironmentParameters) DeepCopy() *BeanstalkEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentSettingObservation) DeepCopyInto(out *BeanstalkEnvironmentSettingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentSettingObservation.
func (in *BeanstalkEnvironmentSettingObservation) DeepCopy() *BeanstalkEnvironmentSettingObservation {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentSettingParameters) DeepCopyInto(out *BeanstalkEnvironmentSettingParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentSettingParameters.
func (in *BeanstalkEnvironmentSettingParameters) DeepCopy() *BeanstalkEnvironmentSettingParameters {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentSpec) DeepCopyInto(out *BeanstalkEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentSpec.
func (in *BeanstalkEnvironmentSpec) DeepCopy() *BeanstalkEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanstalkEnvironmentStatus) DeepCopyInto(out *BeanstalkEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanstalkEnvironmentStatus.
func (in *BeanstalkEnvironmentStatus) DeepCopy() *BeanstalkEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(BeanstalkEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingObservation) DeepCopyInto(out *SettingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingObservation.
func (in *SettingObservation) DeepCopy() *SettingObservation {
	if in == nil {
		return nil
	}
	out := new(SettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingParameters) DeepCopyInto(out *SettingParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingParameters.
func (in *SettingParameters) DeepCopy() *SettingParameters {
	if in == nil {
		return nil
	}
	out := new(SettingParameters)
	in.DeepCopyInto(out)
	return out
}