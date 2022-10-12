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

type LabelObservation struct {
}

type LabelParameters struct {

	// Key of the label.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of the label.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type RuleObservation struct {
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`

	ShowbackCredentialName *string `json:"showbackCredentialName,omitempty" tf:"showback_credential_name,omitempty"`
}

type RuleParameters struct {

	// Set limit of alerts for all projects. Defaults to `0`.
	// +kubebuilder:validation:Optional
	GlobalAlertLimit *float64 `json:"globalAlertLimit,omitempty" tf:"global_alert_limit,omitempty"`

	// The kind of showback rule: `General` (data source is Taikun) or `External` (data source is external, see `showback_credential_id`).
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// Labels linked to this showback rule.
	// +kubebuilder:validation:Optional
	Label []LabelParameters `json:"label,omitempty" tf:"label,omitempty"`

	// The metric name.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// The name of the showback rule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the organization which owns the showback rule.
	// +crossplane:generate:reference:type=github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Billing in CZK per selected unit.
	// +kubebuilder:validation:Required
	Price *float64 `json:"price" tf:"price,omitempty"`

	// Set limit of alerts for one project. Defaults to `0`.
	// +kubebuilder:validation:Optional
	ProjectAlertLimit *float64 `json:"projectAlertLimit,omitempty" tf:"project_alert_limit,omitempty"`

	// ID of the showback credential.
	// +crossplane:generate:reference:type=ShowbackCredential
	// +kubebuilder:validation:Optional
	ShowbackCredentialID *string `json:"showbackCredentialId,omitempty" tf:"showback_credential_id,omitempty"`

	// +kubebuilder:validation:Optional
	ShowbackCredentialIDRef *v1.Reference `json:"showbackCredentialIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ShowbackCredentialIDSelector *v1.Selector `json:"showbackCredentialIdSelector,omitempty" tf:"-"`

	// The type of showback rule: `Count` (calculate package as unit) or `Sum` (calculate per quantity).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rule is the Schema for the Rules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec"`
	Status            RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
