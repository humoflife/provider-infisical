package secret

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_secret", func(r *config.Resource) {
		r.Kind = "Secret"
		r.ShortGroup = "secret"
	})

	p.AddResourceConfigurator("infisical_secret_folder", func(r *config.Resource) {
		r.Kind = "SecretFolder"
		r.ShortGroup = "secret"

		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
	})

	p.AddResourceConfigurator("infisical_secret_import", func(r *config.Resource) {
		r.Kind = "SecretImport"
		r.ShortGroup = "secret"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_tag", func(r *config.Resource) {
		r.Kind = "SecretTag"
		r.ShortGroup = "secret"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_validation_rule_dynamic_secrets", func(r *config.Resource) {
		r.Kind = "SecretValidationRuleDynamicSecrets"
		r.ShortGroup = "secret"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_validation_rule_secret_rotations", func(r *config.Resource) {
		r.Kind = "SecretValidationRuleSecretRotations"
		r.ShortGroup = "secret"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_validation_rule_static_secrets", func(r *config.Resource) {
		r.Kind = "SecretValidationRuleStaticSecrets"
		r.ShortGroup = "secret"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
