package secretrotation

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_secret_rotation_azure_client_secret", func(r *config.Resource) {
		r.Kind = "SecretRotationAzureClientSecret"
		r.ShortGroup = "secretrotation"

		r.References["connection_id"] = config.Reference{
			TerraformName: "infisical_app_connection_azure_client_secrets",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_auth0_client_secret", func(r *config.Resource) {
		r.Kind = "SecretRotationAuth0ClientSecret"
		r.ShortGroup = "secretrotation"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_aws_iam_user_secret", func(r *config.Resource) {
		r.Kind = "SecretRotationAwsIamUserSecret"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_datadog_api_key", func(r *config.Resource) {
		r.Kind = "SecretRotationDatadogApiKey"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_datadog"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_ldap_password", func(r *config.Resource) {
		r.Kind = "SecretRotationLdapPassword"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_ldap"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_litellm_api_key", func(r *config.Resource) {
		r.Kind = "SecretRotationLitellmApiKey"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_litellm"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_mssql_credentials", func(r *config.Resource) {
		r.Kind = "SecretRotationMssqlCredentials"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_mssql"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_mysql_credentials", func(r *config.Resource) {
		r.Kind = "SecretRotationMysqlCredentials"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_mysql"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_oracledb_credentials", func(r *config.Resource) {
		r.Kind = "SecretRotationOracledbCredentials"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_oracledb"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_rotation_postgres_credentials", func(r *config.Resource) {
		r.Kind = "SecretRotationPostgresCredentials"
		r.ShortGroup = "secretrotation"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_postgres"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
