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
	Cni *string `json:"cni,omitempty" tf:"cni,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`
}

type ProfileParameters struct {

	// Whether to expose the Service on each Node's IP at a static port, the NodePort. You'll be able to contact the NodePort Service, from outside the cluster, by requesting `<NodeIP>:<NodePort>`. Defaults to `false`.
	// +kubebuilder:validation:Optional
	BastionProxy *bool `json:"bastionProxy,omitempty" tf:"bastion_proxy,omitempty"`

	// The load-balancing solution: `None`, `Octavia` or `Taikun`. `Octavia` and `Taikun` are only available for OpenStack cloud. Defaults to `Octavia`.
	// +kubebuilder:validation:Optional
	LoadBalancingSolution *string `json:"loadBalancingSolution,omitempty" tf:"load_balancing_solution,omitempty"`

	// Indicates whether to lock the Kubernetes profile. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// The name of the Kubernetes profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the organization which owns the Kubernetes profile.
	// +crossplane:generate:reference:type=github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// When enabled, the workload will also run on master nodes (not recommended). Defaults to `false`.
	// +kubebuilder:validation:Optional
	ScheduleOnMaster *bool `json:"scheduleOnMaster,omitempty" tf:"schedule_on_master,omitempty"`

	// If not enabled, the cluster name will be cluster.local. Defaults to `true`.
	// +kubebuilder:validation:Optional
	UniqueClusterName *bool `json:"uniqueClusterName,omitempty" tf:"unique_cluster_name,omitempty"`
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
