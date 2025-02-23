//go:build !ignore_autogenerated
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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskObservation) DeepCopyInto(out *DiskObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskObservation.
func (in *DiskObservation) DeepCopy() *DiskObservation {
	if in == nil {
		return nil
	}
	out := new(DiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskParameters) DeepCopyInto(out *DiskParameters) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.LunID != nil {
		in, out := &in.LunID, &out.LunID
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskParameters.
func (in *DiskParameters) DeepCopy() *DiskParameters {
	if in == nil {
		return nil
	}
	out := new(DiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNodeLabelObservation) DeepCopyInto(out *KubernetesNodeLabelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNodeLabelObservation.
func (in *KubernetesNodeLabelObservation) DeepCopy() *KubernetesNodeLabelObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesNodeLabelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNodeLabelParameters) DeepCopyInto(out *KubernetesNodeLabelParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNodeLabelParameters.
func (in *KubernetesNodeLabelParameters) DeepCopy() *KubernetesNodeLabelParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesNodeLabelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.AccessIP != nil {
		in, out := &in.AccessIP, &out.AccessIP
		*out = new(string)
		**out = **in
	}
	if in.AlertingProfileName != nil {
		in, out := &in.AlertingProfileName, &out.AlertingProfileName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ServerBastion != nil {
		in, out := &in.ServerBastion, &out.ServerBastion
		*out = make([]ServerBastionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerKubemaster != nil {
		in, out := &in.ServerKubemaster, &out.ServerKubemaster
		*out = make([]ServerKubemasterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerKubeworker != nil {
		in, out := &in.ServerKubeworker, &out.ServerKubeworker
		*out = make([]ServerKubeworkerObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VM != nil {
		in, out := &in.VM, &out.VM
		*out = make([]VMObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.AccessProfileID != nil {
		in, out := &in.AccessProfileID, &out.AccessProfileID
		*out = new(string)
		**out = **in
	}
	if in.AccessProfileIDRef != nil {
		in, out := &in.AccessProfileIDRef, &out.AccessProfileIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.AccessProfileIDSelector != nil {
		in, out := &in.AccessProfileIDSelector, &out.AccessProfileIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AlertingProfileID != nil {
		in, out := &in.AlertingProfileID, &out.AlertingProfileID
		*out = new(string)
		**out = **in
	}
	if in.AlertingProfileIDRef != nil {
		in, out := &in.AlertingProfileIDRef, &out.AlertingProfileIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.AlertingProfileIDSelector != nil {
		in, out := &in.AlertingProfileIDSelector, &out.AlertingProfileIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoUpgrade != nil {
		in, out := &in.AutoUpgrade, &out.AutoUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.BackupCredentialID != nil {
		in, out := &in.BackupCredentialID, &out.BackupCredentialID
		*out = new(string)
		**out = **in
	}
	if in.BackupCredentialIDRef != nil {
		in, out := &in.BackupCredentialIDRef, &out.BackupCredentialIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.BackupCredentialIDSelector != nil {
		in, out := &in.BackupCredentialIDSelector, &out.BackupCredentialIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudCredentialID != nil {
		in, out := &in.CloudCredentialID, &out.CloudCredentialID
		*out = new(string)
		**out = **in
	}
	if in.CloudCredentialIDRef != nil {
		in, out := &in.CloudCredentialIDRef, &out.CloudCredentialIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.CloudCredentialIDSelector != nil {
		in, out := &in.CloudCredentialIDSelector, &out.CloudCredentialIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DeleteOnExpiration != nil {
		in, out := &in.DeleteOnExpiration, &out.DeleteOnExpiration
		*out = new(bool)
		**out = **in
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KubernetesProfileID != nil {
		in, out := &in.KubernetesProfileID, &out.KubernetesProfileID
		*out = new(string)
		**out = **in
	}
	if in.KubernetesProfileIDRef != nil {
		in, out := &in.KubernetesProfileIDRef, &out.KubernetesProfileIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KubernetesProfileIDSelector != nil {
		in, out := &in.KubernetesProfileIDSelector, &out.KubernetesProfileIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Lock != nil {
		in, out := &in.Lock, &out.Lock
		*out = new(bool)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationIDRef != nil {
		in, out := &in.OrganizationIDRef, &out.OrganizationIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.OrganizationIDSelector != nil {
		in, out := &in.OrganizationIDSelector, &out.OrganizationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyProfileID != nil {
		in, out := &in.PolicyProfileID, &out.PolicyProfileID
		*out = new(string)
		**out = **in
	}
	if in.PolicyProfileIDRef != nil {
		in, out := &in.PolicyProfileIDRef, &out.PolicyProfileIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.PolicyProfileIDSelector != nil {
		in, out := &in.PolicyProfileIDSelector, &out.PolicyProfileIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.QuotaCPUUnits != nil {
		in, out := &in.QuotaCPUUnits, &out.QuotaCPUUnits
		*out = new(float64)
		**out = **in
	}
	if in.QuotaDiskSize != nil {
		in, out := &in.QuotaDiskSize, &out.QuotaDiskSize
		*out = new(float64)
		**out = **in
	}
	if in.QuotaRAMSize != nil {
		in, out := &in.QuotaRAMSize, &out.QuotaRAMSize
		*out = new(float64)
		**out = **in
	}
	if in.QuotaVMCPUUnits != nil {
		in, out := &in.QuotaVMCPUUnits, &out.QuotaVMCPUUnits
		*out = new(float64)
		**out = **in
	}
	if in.QuotaVMRAMSize != nil {
		in, out := &in.QuotaVMRAMSize, &out.QuotaVMRAMSize
		*out = new(float64)
		**out = **in
	}
	if in.QuotaVMVolumeSize != nil {
		in, out := &in.QuotaVMVolumeSize, &out.QuotaVMVolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.RouterIDEndRange != nil {
		in, out := &in.RouterIDEndRange, &out.RouterIDEndRange
		*out = new(float64)
		**out = **in
	}
	if in.RouterIDStartRange != nil {
		in, out := &in.RouterIDStartRange, &out.RouterIDStartRange
		*out = new(float64)
		**out = **in
	}
	if in.ServerBastion != nil {
		in, out := &in.ServerBastion, &out.ServerBastion
		*out = make([]ServerBastionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerKubemaster != nil {
		in, out := &in.ServerKubemaster, &out.ServerKubemaster
		*out = make([]ServerKubemasterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerKubeworker != nil {
		in, out := &in.ServerKubeworker, &out.ServerKubeworker
		*out = make([]ServerKubeworkerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TaikunLBFlavor != nil {
		in, out := &in.TaikunLBFlavor, &out.TaikunLBFlavor
		*out = new(string)
		**out = **in
	}
	if in.VM != nil {
		in, out := &in.VM, &out.VM
		*out = make([]VMParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerBastionObservation) DeepCopyInto(out *ServerBastionObservation) {
	*out = *in
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerBastionObservation.
func (in *ServerBastionObservation) DeepCopy() *ServerBastionObservation {
	if in == nil {
		return nil
	}
	out := new(ServerBastionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerBastionParameters) DeepCopyInto(out *ServerBastionParameters) {
	*out = *in
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerBastionParameters.
func (in *ServerBastionParameters) DeepCopy() *ServerBastionParameters {
	if in == nil {
		return nil
	}
	out := new(ServerBastionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubemasterObservation) DeepCopyInto(out *ServerKubemasterObservation) {
	*out = *in
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubemasterObservation.
func (in *ServerKubemasterObservation) DeepCopy() *ServerKubemasterObservation {
	if in == nil {
		return nil
	}
	out := new(ServerKubemasterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubemasterParameters) DeepCopyInto(out *ServerKubemasterParameters) {
	*out = *in
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(string)
		**out = **in
	}
	if in.KubernetesNodeLabel != nil {
		in, out := &in.KubernetesNodeLabel, &out.KubernetesNodeLabel
		*out = make([]KubernetesNodeLabelParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubemasterParameters.
func (in *ServerKubemasterParameters) DeepCopy() *ServerKubemasterParameters {
	if in == nil {
		return nil
	}
	out := new(ServerKubemasterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubeworkerKubernetesNodeLabelObservation) DeepCopyInto(out *ServerKubeworkerKubernetesNodeLabelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubeworkerKubernetesNodeLabelObservation.
func (in *ServerKubeworkerKubernetesNodeLabelObservation) DeepCopy() *ServerKubeworkerKubernetesNodeLabelObservation {
	if in == nil {
		return nil
	}
	out := new(ServerKubeworkerKubernetesNodeLabelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubeworkerKubernetesNodeLabelParameters) DeepCopyInto(out *ServerKubeworkerKubernetesNodeLabelParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubeworkerKubernetesNodeLabelParameters.
func (in *ServerKubeworkerKubernetesNodeLabelParameters) DeepCopy() *ServerKubeworkerKubernetesNodeLabelParameters {
	if in == nil {
		return nil
	}
	out := new(ServerKubeworkerKubernetesNodeLabelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubeworkerObservation) DeepCopyInto(out *ServerKubeworkerObservation) {
	*out = *in
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubeworkerObservation.
func (in *ServerKubeworkerObservation) DeepCopy() *ServerKubeworkerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerKubeworkerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerKubeworkerParameters) DeepCopyInto(out *ServerKubeworkerParameters) {
	*out = *in
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(string)
		**out = **in
	}
	if in.KubernetesNodeLabel != nil {
		in, out := &in.KubernetesNodeLabel, &out.KubernetesNodeLabel
		*out = make([]ServerKubeworkerKubernetesNodeLabelParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerKubeworkerParameters.
func (in *ServerKubeworkerParameters) DeepCopy() *ServerKubeworkerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerKubeworkerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagObservation) DeepCopyInto(out *TagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagObservation.
func (in *TagObservation) DeepCopy() *TagObservation {
	if in == nil {
		return nil
	}
	out := new(TagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagParameters) DeepCopyInto(out *TagParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagParameters.
func (in *TagParameters) DeepCopy() *TagParameters {
	if in == nil {
		return nil
	}
	out := new(TagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMObservation) DeepCopyInto(out *VMObservation) {
	*out = *in
	if in.AccessIP != nil {
		in, out := &in.AccessIP, &out.AccessIP
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMObservation.
func (in *VMObservation) DeepCopy() *VMObservation {
	if in == nil {
		return nil
	}
	out := new(VMObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMParameters) DeepCopyInto(out *VMParameters) {
	*out = *in
	if in.CloudInit != nil {
		in, out := &in.CloudInit, &out.CloudInit
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(bool)
		**out = **in
	}
	if in.StandaloneProfileID != nil {
		in, out := &in.StandaloneProfileID, &out.StandaloneProfileID
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMParameters.
func (in *VMParameters) DeepCopy() *VMParameters {
	if in == nil {
		return nil
	}
	out := new(VMParameters)
	in.DeepCopyInto(out)
	return out
}
