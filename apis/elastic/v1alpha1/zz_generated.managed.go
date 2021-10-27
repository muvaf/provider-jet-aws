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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this BeanstalkApplication.
func (mg *BeanstalkApplication) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BeanstalkApplication.
func (mg *BeanstalkApplication) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BeanstalkApplication.
func (mg *BeanstalkApplication) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BeanstalkApplication.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BeanstalkApplication) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BeanstalkApplication.
func (mg *BeanstalkApplication) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BeanstalkApplication.
func (mg *BeanstalkApplication) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BeanstalkApplication.
func (mg *BeanstalkApplication) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BeanstalkApplication.
func (mg *BeanstalkApplication) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BeanstalkApplication.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BeanstalkApplication) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BeanstalkApplication.
func (mg *BeanstalkApplication) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BeanstalkApplicationVersion.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BeanstalkApplicationVersion) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BeanstalkApplicationVersion.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BeanstalkApplicationVersion) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BeanstalkApplicationVersion.
func (mg *BeanstalkApplicationVersion) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BeanstalkConfigurationTemplate.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BeanstalkConfigurationTemplate) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BeanstalkConfigurationTemplate.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BeanstalkConfigurationTemplate) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BeanstalkConfigurationTemplate.
func (mg *BeanstalkConfigurationTemplate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BeanstalkEnvironment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BeanstalkEnvironment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BeanstalkEnvironment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BeanstalkEnvironment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BeanstalkEnvironment.
func (mg *BeanstalkEnvironment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}