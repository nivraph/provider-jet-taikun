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
}
