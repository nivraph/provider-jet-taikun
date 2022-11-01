package resources

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("taikun_organization", func(r *config.Resource) {
		r.ShortGroup = "organization"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_access_profile", func(r *config.Resource) {
		r.ShortGroup = "accessProfile"
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_slack_configuration", func(r *config.Resource) {
		r.ShortGroup = "slackConfiguration"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_alerting_profile", func(r *config.Resource) {
		r.ShortGroup = "alertingProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["slack_configuration_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/slackconfiguration/v1alpha1.Configuration",
		}
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_kubernetes_profile", func(r *config.Resource) {
		r.ShortGroup = "kubernetesProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_policy_profile", func(r *config.Resource) {
		r.ShortGroup = "policyProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential", func(r *config.Resource) {
		r.ShortGroup = "cloudCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_aws", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialAws"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_openstack", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialOpenstack"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_showback_credential", func(r *config.Resource) {
		r.ShortGroup = "showbackCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_showback_rule", func(r *config.Resource) {
		r.ShortGroup = "showbackRule"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.References["showback_credential_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/showbackcredential/v1alpha1.Credential",
		}
	})
	p.AddResourceConfigurator("taikun_billing_credential", func(r *config.Resource) {
		r.ShortGroup = "billingCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_billing_rule", func(r *config.Resource) {
		r.ShortGroup = "billingRule"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.References["billing_credential_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/billingcredential/v1alpha1.Credential",
		}
	})
	p.AddResourceConfigurator("taikun_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_standalone_profile", func(r *config.Resource) {
		r.ShortGroup = "standaloneProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_backup_policy", func(r *config.Resource) {
		r.ShortGroup = "backupPolicy"
		r.ExternalName = config.IdentifierFromProvider
		r.References["project_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/project/v1alpha1.Project",
		}
	})
	p.AddResourceConfigurator("taikun_backup_credential", func(r *config.Resource) {
		r.ShortGroup = "backupCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_gcp", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialGcp"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.References["billing_account_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/billingcredential/v1alpha1.Credential",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_azure", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialAzure"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
	})
	p.AddResourceConfigurator("taikun_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.References["access_profile_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/accessprofile/v1alpha1.Profile",
		}
		r.References["alerting_profile_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/alertingprofile/v1alpha1.Profile",
		}
		r.References["kubernetes_profile_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/kubernetesprofile/v1alpha1.Profile",
		}
		r.References["policy_profile_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/policyprofile/v1alpha1.Profile",
		}
		r.References["cloud_credential_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/cloudcredentialopenstack/v1alpha1.CredentialOpenstack",
		}
		r.References["standalone_profile_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/standaloneprofile/v1alpha1.Profile",
		}
		r.References["backup_credential_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/backupcredential/v1alpha1.Credential",
		}
	})
	p.AddResourceConfigurator("taikun_organization_billing_rule_attachment", func(r *config.Resource) {
		r.ShortGroup = "organizationBillingRuleAttachment"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/organization/v1alpha1.Organization",
		}
		r.References["billing_rule_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/billingrule/v1alpha1.Rule",
		}
	})
	p.AddResourceConfigurator("taikun_kubeconfig", func(r *config.Resource) {
		r.ShortGroup = "kubeconfig"
		r.ExternalName = config.IdentifierFromProvider
		r.References["project_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/project/v1alpha1.Project",
		}
		r.References["user_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/user/v1alpha1.User",
		}
	})
	p.AddResourceConfigurator("taikun_project_user_attachment", func(r *config.Resource) {
		r.ShortGroup = "projectUserAttachment"
		r.ExternalName = config.IdentifierFromProvider
		r.References["project_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/project/v1alpha1.Project",
		}
		r.References["user_id"] = config.Reference{
			Type: "github.com/nivraph/provider-jet-taikun/apis/user/v1alpha1.User",
		}
	})
}
