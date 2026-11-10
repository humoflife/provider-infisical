package kms

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the kms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_external_kms_aws", func(r *config.Resource) {
		r.Kind = "ExternalAws"
		r.ShortGroup = "kms"
	})

	p.AddResourceConfigurator("infisical_kms_key", func(r *config.Resource) {
		r.Kind = "Key"
		r.ShortGroup = "kms"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
