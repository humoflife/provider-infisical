package organization

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the organization group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_alert", func(r *config.Resource) {
		r.Kind = "Alert"
		r.ShortGroup = "organization"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_org_role", func(r *config.Resource) {
		r.Kind = "Role"
		r.ShortGroup = "organization"
	})

	p.AddResourceConfigurator("infisical_sub_organization", func(r *config.Resource) {
		r.Kind = "SubOrganization"
		r.ShortGroup = "organization"
	})

	p.AddResourceConfigurator("infisical_sub_organization_group", func(r *config.Resource) {
		r.Kind = "SubOrganizationGroup"
		r.ShortGroup = "organization"
		r.References["group_id"] = config.Reference{TerraformName: "infisical_group"}
	})
}
