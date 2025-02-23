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

type ProfileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`
}

type ProfileParameters struct {

	// Requires container images to begin with a string from the specified list.
	// +kubebuilder:validation:Optional
	AllowedRepos []*string `json:"allowedRepos,omitempty" tf:"allowed_repos,omitempty"`

	// Requires Ingress resources to be HTTPS only. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ForbidHTTPIngress *bool `json:"forbidHttpIngress,omitempty" tf:"forbid_http_ingress,omitempty"`

	// Disallows all Services with type NodePort. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ForbidNodePort *bool `json:"forbidNodePort,omitempty" tf:"forbid_node_port,omitempty"`

	// Container images must have an image tag different from the ones in the list.
	// +kubebuilder:validation:Optional
	ForbiddenTags []*string `json:"forbiddenTags,omitempty" tf:"forbidden_tags,omitempty"`

	// List of allowed Ingress rule hosts.
	// +kubebuilder:validation:Optional
	IngressWhitelist []*string `json:"ingressWhitelist,omitempty" tf:"ingress_whitelist,omitempty"`

	// Indicates whether to lock the Policy profile. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// The name of the Policy profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the organization which owns the Policy profile.
	// +crossplane:generate:reference:type=github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Requires Pods to have readiness and liveness probes. Defaults to `false`.
	// +kubebuilder:validation:Optional
	RequireProbe *bool `json:"requireProbe,omitempty" tf:"require_probe,omitempty"`

	// Requires all Ingress rule hosts to be unique. Defaults to `false`.
	// +kubebuilder:validation:Optional
	UniqueIngress *bool `json:"uniqueIngress,omitempty" tf:"unique_ingress,omitempty"`

	// Whether services must have globally unique service selectors or not. Defaults to `false`.
	// +kubebuilder:validation:Optional
	UniqueServiceSelector *bool `json:"uniqueServiceSelector,omitempty" tf:"unique_service_selector,omitempty"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileParameters `json:"forProvider"`
}

// ProfileStatus defines the observed state of Profile.
type ProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Profile is the Schema for the Profiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProfileSpec   `json:"spec"`
	Status            ProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileList contains a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Repository type metadata.
var (
	Profile_Kind             = "Profile"
	Profile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Profile_Kind}.String()
	Profile_KindAPIVersion   = Profile_Kind + "." + CRDGroupVersion.String()
	Profile_GroupVersionKind = CRDGroupVersion.WithKind(Profile_Kind)
)

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
