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
	credential "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/backupcredential/credential"
	policy "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/backuppolicy/policy"
	credentialbillingcredential "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/billingcredential/credential"
	rule "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/billingrule/rule"
	credentialaws "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialaws/credentialaws"
	credentialazure "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialazure/credentialazure"
	credentialgcp "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialgcp/credentialgcp"
	credentialopenstack "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/cloudcredentialopenstack/credentialopenstack"
	profilekubernetesprofile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/kubernetesprofile/profile"
	organization "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/organization/organization"
	billingruleattachment "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/organizationbillingruleattachment/billingruleattachment"
	profilepolicyprofile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/policyprofile/profile"
	project "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/project/project"
	providerconfig "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/providerconfig"
	credentialshowbackcredential "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/showbackcredential/credential"
	ruleshowbackrule "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/showbackrule/rule"
	configuration "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/slackconfiguration/configuration"
	profilestandaloneprofile "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/standaloneprofile/profile"
	user "github.com/crossplane-contrib/provider-jet-taikun/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		profile.Setup,
		profilealertingprofile.Setup,
		credential.Setup,
		policy.Setup,
		credentialbillingcredential.Setup,
		rule.Setup,
		credentialaws.Setup,
		credentialazure.Setup,
		credentialgcp.Setup,
		credentialopenstack.Setup,
		profilekubernetesprofile.Setup,
		organization.Setup,
		billingruleattachment.Setup,
		profilepolicyprofile.Setup,
		project.Setup,
		providerconfig.Setup,
		credentialshowbackcredential.Setup,
		ruleshowbackrule.Setup,
		configuration.Setup,
		profilestandaloneprofile.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
