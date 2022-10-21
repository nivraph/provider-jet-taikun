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

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/nivraph/provider-jet-taikun/apis/accessprofile/v1alpha1"
	v1alpha11 "github.com/nivraph/provider-jet-taikun/apis/alertingprofile/v1alpha1"
	v1alpha12 "github.com/nivraph/provider-jet-taikun/apis/backupcredential/v1alpha1"
	v1alpha13 "github.com/nivraph/provider-jet-taikun/apis/cloudcredentialopenstack/v1alpha1"
	v1alpha14 "github.com/nivraph/provider-jet-taikun/apis/kubernetesprofile/v1alpha1"
	v1alpha15 "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1"
	v1alpha16 "github.com/nivraph/provider-jet-taikun/apis/policyprofile/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Project.
func (mg *Project) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessProfileID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccessProfileIDRef,
		Selector:     mg.Spec.ForProvider.AccessProfileIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProfileList{},
			Managed: &v1alpha1.Profile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessProfileID")
	}
	mg.Spec.ForProvider.AccessProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessProfileIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AlertingProfileID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AlertingProfileIDRef,
		Selector:     mg.Spec.ForProvider.AlertingProfileIDSelector,
		To: reference.To{
			List:    &v1alpha11.ProfileList{},
			Managed: &v1alpha11.Profile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AlertingProfileID")
	}
	mg.Spec.ForProvider.AlertingProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AlertingProfileIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupCredentialID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackupCredentialIDRef,
		Selector:     mg.Spec.ForProvider.BackupCredentialIDSelector,
		To: reference.To{
			List:    &v1alpha12.CredentialList{},
			Managed: &v1alpha12.Credential{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupCredentialID")
	}
	mg.Spec.ForProvider.BackupCredentialID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupCredentialIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudCredentialID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudCredentialIDRef,
		Selector:     mg.Spec.ForProvider.CloudCredentialIDSelector,
		To: reference.To{
			List:    &v1alpha13.CredentialOpenstackList{},
			Managed: &v1alpha13.CredentialOpenstack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudCredentialID")
	}
	mg.Spec.ForProvider.CloudCredentialID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudCredentialIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KubernetesProfileID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KubernetesProfileIDRef,
		Selector:     mg.Spec.ForProvider.KubernetesProfileIDSelector,
		To: reference.To{
			List:    &v1alpha14.ProfileList{},
			Managed: &v1alpha14.Profile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KubernetesProfileID")
	}
	mg.Spec.ForProvider.KubernetesProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KubernetesProfileIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrganizationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationIDRef,
		Selector:     mg.Spec.ForProvider.OrganizationIDSelector,
		To: reference.To{
			List:    &v1alpha15.OrganizationList{},
			Managed: &v1alpha15.Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrganizationID")
	}
	mg.Spec.ForProvider.OrganizationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyProfileID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyProfileIDRef,
		Selector:     mg.Spec.ForProvider.PolicyProfileIDSelector,
		To: reference.To{
			List:    &v1alpha16.ProfileList{},
			Managed: &v1alpha16.Profile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyProfileID")
	}
	mg.Spec.ForProvider.PolicyProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyProfileIDRef = rsp.ResolvedReference

	return nil
}
