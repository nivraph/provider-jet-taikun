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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	profile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/accessprofile/profile"
	profilealertingprofile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/alertingprofile/profile"
	credentialaws "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialaws/credentialaws"
	credentialopenstack "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialopenstack/credentialopenstack"
	profilekubernetesprofile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/kubernetesprofile/profile"
	organization "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/organization/organization"
	providerconfig "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/providerconfig"
	configuration "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/slackconfiguration/configuration"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		profile.Setup,
		profilealertingprofile.Setup,
		credentialaws.Setup,
		credentialopenstack.Setup,
		profilekubernetesprofile.Setup,
		organization.Setup,
		providerconfig.Setup,
		configuration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
