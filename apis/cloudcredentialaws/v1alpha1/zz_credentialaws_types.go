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

type CredentialAwsObservation struct {
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`
}

type CredentialAwsParameters struct {

	// The AWS access key ID.
	// +kubebuilder:validation:Required
	AccessKeyIDSecretRef v1.SecretKeySelector `json:"accessKeyIdSecretRef" tf:"-"`

	// The AWS availability zone for the region.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Indicates whether to lock the AWS cloud credential. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// The name of the AWS cloud credential.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the organization which owns the AWS cloud credential.
	// +crossplane:generate:reference:type=github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// The AWS region.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The AWS secret access key.
	// +kubebuilder:validation:Required
	SecretAccessKeySecretRef v1.SecretKeySelector `json:"secretAccessKeySecretRef" tf:"-"`
}

// CredentialAwsSpec defines the desired state of CredentialAws
type CredentialAwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialAwsParameters `json:"forProvider"`
}

// CredentialAwsStatus defines the observed state of CredentialAws.
type CredentialAwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialAwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialAws is the Schema for the CredentialAwss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type CredentialAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialAwsSpec   `json:"spec"`
	Status            CredentialAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialAwsList contains a list of CredentialAwss
type CredentialAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CredentialAws `json:"items"`
}

// Repository type metadata.
var (
	CredentialAws_Kind             = "CredentialAws"
	CredentialAws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CredentialAws_Kind}.String()
	CredentialAws_KindAPIVersion   = CredentialAws_Kind + "." + CRDGroupVersion.String()
	CredentialAws_GroupVersionKind = CRDGroupVersion.WithKind(CredentialAws_Kind)
)

func init() {
	SchemeBuilder.Register(&CredentialAws{}, &CredentialAwsList{})
}
