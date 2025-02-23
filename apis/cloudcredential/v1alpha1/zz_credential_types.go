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

type CredentialObservation struct {
	BillingAccountName *string `json:"billingAccountName,omitempty" tf:"billing_account_name,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CredentialParameters struct {

	// The AWS access key ID. Required for AWS.
	// +kubebuilder:validation:Optional
	AccessKeyIDSecretRef *v1.SecretKeySelector `json:"accessKeyIdSecretRef,omitempty" tf:"-"`

	// The availability zone of the cloud credential. Optional for Openstack. Required for AWS and Azure. See `zone` for GCP.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID of the GCP credential's billing account. Conflicts with: `import_project`.
	// +kubebuilder:validation:Optional
	BillingAccountID *string `json:"billingAccountId,omitempty" tf:"billing_account_id,omitempty"`

	// The Azure client ID. Required for Azure.
	// +kubebuilder:validation:Optional
	ClientIDSecretRef *v1.SecretKeySelector `json:"clientIdSecretRef,omitempty" tf:"-"`

	// The Azure client secret. Required for Azure.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// The path of the GCP credential's configuration file. Required for GCP.
	// +kubebuilder:validation:Optional
	ConfigFile *string `json:"configFile,omitempty" tf:"config_file,omitempty"`

	// The OpenStack domain. Required for Openstack.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The folder ID of the GCP credential. Conflicts with: `import_project`.
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
	// +kubebuilder:validation:Optional
	ImportProject *bool `json:"importProject,omitempty" tf:"import_project,omitempty"`

	// The OpenStack network subnet ID to import a network.
	// +kubebuilder:validation:Optional
	ImportedNetworkSubnetID *string `json:"importedNetworkSubnetId,omitempty" tf:"imported_network_subnet_id,omitempty"`

	// The Azure location. Required for Azure.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Indicates whether to lock the cloud credential. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// The name of the cloud credential.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the organization which owns the cloud credential.
	// +crossplane:generate:reference:type=github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// The OpenStack password. Required for Openstack.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The OpenStack project name. Required for Openstack.
	// +kubebuilder:validation:Optional
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// The name of the public OpenStack network to use. Required for Openstack.
	// +kubebuilder:validation:Optional
	PublicNetworkName *string `json:"publicNetworkName,omitempty" tf:"public_network_name,omitempty"`

	// The region of the cloud credential. Required for Openstack, AWS and GCP.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The AWS secret access key. Required for AWS.
	// +kubebuilder:validation:Optional
	SecretAccessKeySecretRef *v1.SecretKeySelector `json:"secretAccessKeySecretRef,omitempty" tf:"-"`

	// The Azure subscription ID. Required for Azure.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Azure tenant ID. Required for Azure.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of the cloud credential.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The OpenStack authentication URL. Required for Openstack.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The OpenStack user. Required for Openstack.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// The OpenStack type of volume.
	// +kubebuilder:validation:Optional
	VolumeTypeName *string `json:"volumeTypeName,omitempty" tf:"volume_type_name,omitempty"`

	// The zone of the GCP credential. Required for GCP.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// CredentialSpec defines the desired state of Credential
type CredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialParameters `json:"forProvider"`
}

// CredentialStatus defines the observed state of Credential.
type CredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Credential is the Schema for the Credentials API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialSpec   `json:"spec"`
	Status            CredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialList contains a list of Credentials
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items"`
}

// Repository type metadata.
var (
	Credential_Kind             = "Credential"
	Credential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Credential_Kind}.String()
	Credential_KindAPIVersion   = Credential_Kind + "." + CRDGroupVersion.String()
	Credential_GroupVersionKind = CRDGroupVersion.WithKind(Credential_Kind)
)

func init() {
	SchemeBuilder.Register(&Credential{}, &CredentialList{})
}
