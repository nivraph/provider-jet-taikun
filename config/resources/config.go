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
}
