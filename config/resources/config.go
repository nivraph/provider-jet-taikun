package organization

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("taikun_organization", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "organization"
	})
	p.AddResourceConfigurator("taikun_access_profile", func(r *config.Resource) {
		r.ShortGroup = "accessProfile"
		/*
			r.References["organization_id"] = config.Reference{
				Type: "Organization",
			}
			r.ExternalName = config.IdentifierFromProvider
		*/
	})
}
