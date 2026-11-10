package integration

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the integration group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_integration_aws_parameter_store", func(r *config.Resource) {
		r.Kind = "AwsParameterStore"
		r.ShortGroup = "integration"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_integration_aws_secrets_manager", func(r *config.Resource) {
		r.Kind = "AwsSecretsManager"
		r.ShortGroup = "integration"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_integration_circleci", func(r *config.Resource) {
		r.Kind = "Circleci"
		r.ShortGroup = "integration"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_integration_databricks", func(r *config.Resource) {
		r.Kind = "Databricks"
		r.ShortGroup = "integration"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_integration_gcp_secret_manager", func(r *config.Resource) {
		r.Kind = "GcpSecretManager"
		r.ShortGroup = "integration"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
