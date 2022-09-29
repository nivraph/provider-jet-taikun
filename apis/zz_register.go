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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-taikun/apis/accessprofile/v1alpha1"
	v1alpha1alertingprofile "github.com/crossplane-contrib/provider-jet-taikun/apis/alertingprofile/v1alpha1"
	v1alpha1cloudcredentialaws "github.com/crossplane-contrib/provider-jet-taikun/apis/cloudcredentialaws/v1alpha1"
	v1alpha1cloudcredentialopenstack "github.com/crossplane-contrib/provider-jet-taikun/apis/cloudcredentialopenstack/v1alpha1"
	v1alpha1kubernetesprofile "github.com/crossplane-contrib/provider-jet-taikun/apis/kubernetesprofile/v1alpha1"
	v1alpha1organization "github.com/crossplane-contrib/provider-jet-taikun/apis/organization/v1alpha1"
	v1alpha1policyprofile "github.com/crossplane-contrib/provider-jet-taikun/apis/policyprofile/v1alpha1"
	v1alpha1slackconfiguration "github.com/crossplane-contrib/provider-jet-taikun/apis/slackconfiguration/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-taikun/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1alertingprofile.SchemeBuilder.AddToScheme,
		v1alpha1cloudcredentialaws.SchemeBuilder.AddToScheme,
		v1alpha1cloudcredentialopenstack.SchemeBuilder.AddToScheme,
		v1alpha1kubernetesprofile.SchemeBuilder.AddToScheme,
		v1alpha1organization.SchemeBuilder.AddToScheme,
		v1alpha1policyprofile.SchemeBuilder.AddToScheme,
		v1alpha1slackconfiguration.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
