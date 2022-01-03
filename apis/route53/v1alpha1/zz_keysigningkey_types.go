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

type KeySigningKeyObservation struct {
	DigestAlgorithmMnemonic *string `json:"digestAlgorithmMnemonic,omitempty" tf:"digest_algorithm_mnemonic,omitempty"`

	DigestAlgorithmType *int64 `json:"digestAlgorithmType,omitempty" tf:"digest_algorithm_type,omitempty"`

	DigestValue *string `json:"digestValue,omitempty" tf:"digest_value,omitempty"`

	DnskeyRecord *string `json:"dnskeyRecord,omitempty" tf:"dnskey_record,omitempty"`

	DsRecord *string `json:"dsRecord,omitempty" tf:"ds_record,omitempty"`

	Flag *int64 `json:"flag,omitempty" tf:"flag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyTag *int64 `json:"keyTag,omitempty" tf:"key_tag,omitempty"`

	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	SigningAlgorithmMnemonic *string `json:"signingAlgorithmMnemonic,omitempty" tf:"signing_algorithm_mnemonic,omitempty"`

	SigningAlgorithmType *int64 `json:"signingAlgorithmType,omitempty" tf:"signing_algorithm_type,omitempty"`
}

type KeySigningKeyParameters struct {

	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	HostedZoneIDRef *v1.Reference `json:"hostedZoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostedZoneIDSelector *v1.Selector `json:"hostedZoneIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1.KMSKeyARN()
	// +kubebuilder:validation:Optional
	KeyManagementServiceArn *string `json:"keyManagementServiceArn,omitempty" tf:"key_management_service_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KeyManagementServiceArnRef *v1.Reference `json:"keyManagementServiceArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyManagementServiceArnSelector *v1.Selector `json:"keyManagementServiceArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// KeySigningKeySpec defines the desired state of KeySigningKey
type KeySigningKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeySigningKeyParameters `json:"forProvider"`
}

// KeySigningKeyStatus defines the observed state of KeySigningKey.
type KeySigningKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeySigningKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeySigningKey is the Schema for the KeySigningKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type KeySigningKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySigningKeySpec   `json:"spec"`
	Status            KeySigningKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeySigningKeyList contains a list of KeySigningKeys
type KeySigningKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeySigningKey `json:"items"`
}

// Repository type metadata.
var (
	KeySigningKey_Kind             = "KeySigningKey"
	KeySigningKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeySigningKey_Kind}.String()
	KeySigningKey_KindAPIVersion   = KeySigningKey_Kind + "." + CRDGroupVersion.String()
	KeySigningKey_GroupVersionKind = CRDGroupVersion.WithKind(KeySigningKey_Kind)
)

func init() {
	SchemeBuilder.Register(&KeySigningKey{}, &KeySigningKeyList{})
}