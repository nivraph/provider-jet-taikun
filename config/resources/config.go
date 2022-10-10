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
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_slack_configuration", func(r *config.Resource) {
		r.ShortGroup = "slackConfiguration"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_alerting_profile", func(r *config.Resource) {
		r.ShortGroup = "alertingProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["slackConfigurationId"] = config.Reference{
			Type: "SlackConfiguration",
		}
	})
	p.AddResourceConfigurator("taikun_kubernetes_profile", func(r *config.Resource) {
		r.ShortGroup = "kubernetesProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_policy_profile", func(r *config.Resource) {
		r.ShortGroup = "policyProfile"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_aws", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialAws"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_openstack", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialOpenstack"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_showback_credential", func(r *config.Resource) {
		r.ShortGroup = "showbackCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_showback_rule", func(r *config.Resource) {
		r.ShortGroup = "showbackRule"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.References["showbackCredentialId"] = config.Reference{
			Type: "ShowbackCredential",
		}
	})
	p.AddResourceConfigurator("taikun_billing_credential", func(r *config.Resource) {
		r.ShortGroup = "billingCredential"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_billing_rule", func(r *config.Resource) {
		r.ShortGroup = "billingRule"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.References["billingCredentialId"] = config.Reference{
			Type: "BillingCredential",
		}
	})
	p.AddResourceConfigurator("taikun_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_standalone_profile", func(r *config.Resource) {
		r.ShortGroup = "standaloneProfile"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_backup_policy", func(r *config.Resource) {
		r.ShortGroup = "backupPolicy"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_backup_credential", func(r *config.Resource) {
		r.ShortGroup = "backupCredential"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("taikun_cloud_credential_gcp", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialGcp"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.References["billingAccountId"] = config.Reference{
			Type: "BillingCredential",
		}
	})
	p.AddResourceConfigurator("taikun_cloud_credential_azure", func(r *config.Resource) {
		r.ShortGroup = "cloudCredentialAzure"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
	})
	p.AddResourceConfigurator("taikun_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.References["accessProfileId"] = config.Reference{
			Type: "AccessProfile",
		}
		r.References["alertingProfileId"] = config.Reference{
			Type: "AlertingProfile",
		}
		r.References["KubernetesProfileId"] = config.Reference{
			Type: "KubernetesProfile",
		}
		r.References["policyProfileId"] = config.Reference{
			Type: "PolicyProfile",
		}
		// TODO: add backup credential references
		/*
			r.References["backupCredentialId"] = config.Reference{
				Type: "BackupCredential",
			}
		*/
	})
	p.AddResourceConfigurator("taikun_organization_billing_rule_attachment", func(r *config.Resource) {
		r.ShortGroup = "organizationBillingRuleAttachment"
		r.ExternalName = config.IdentifierFromProvider
		r.References["organizationId"] = config.Reference{
			Type: "Organization",
		}
		r.References["billingRuleId"] = config.Reference{
			Type: "BillingRule",
		}
	})
        p.AddResourceConfigurator("taikun_kubeconfig", func(r *config.Resource) {
                r.ShortGroup = "kubeconfig"
                r.ExternalName = config.IdentifierFromProvider
                r.References["projectId"] = config.Reference{
                        Type: "Project",
                }
        })
        p.AddResourceConfigurator("taikun_project_user_attachment", func(r *config.Resource) {
                r.ShortGroup = "projectUserAttachment"
                r.ExternalName = config.IdentifierFromProvider
                r.References["projectId"] = config.Reference{
                        Type: "Project",
                }
                r.References["userId"] = config.Reference{
                        Type: "User",
                }
        })
}
