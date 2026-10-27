package secretsync

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_secret_sync_github", func(r *config.Resource) {
		r.Kind = "SecretSyncGithub"
		r.ShortGroup = "secretsync" // lowercase not allowed

		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
	})

	p.AddResourceConfigurator("infisical_secret_sync_1password", func(r *config.Resource) {
		r.Kind = "SecretSyncOnePassword"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_1password"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_aws_parameter_store", func(r *config.Resource) {
		r.Kind = "SecretSyncAwsParameterStore"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_aws_secrets_manager", func(r *config.Resource) {
		r.Kind = "SecretSyncAwsSecretsManager"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_azure_app_configuration", func(r *config.Resource) {
		r.Kind = "SecretSyncAzureAppConfiguration"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_azure_app_configuration"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_azure_devops", func(r *config.Resource) {
		r.Kind = "SecretSyncAzureDevops"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_azure_devops"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_azure_key_vault", func(r *config.Resource) {
		r.Kind = "SecretSyncAzureKeyVault"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_azure_key_vault"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_bitbucket", func(r *config.Resource) {
		r.Kind = "SecretSyncBitbucket"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_bitbucket"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_circleci", func(r *config.Resource) {
		r.Kind = "SecretSyncCircleci"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_circleci"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_cloudflare_pages", func(r *config.Resource) {
		r.Kind = "SecretSyncCloudflarePages"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_cloudflare"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_cloudflare_workers", func(r *config.Resource) {
		r.Kind = "SecretSyncCloudflareWorkers"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_cloudflare"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_databricks", func(r *config.Resource) {
		r.Kind = "SecretSyncDatabricks"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_databricks"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_flyio", func(r *config.Resource) {
		r.Kind = "SecretSyncFlyio"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_flyio"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_gcp_secret_manager", func(r *config.Resource) {
		r.Kind = "SecretSyncGcpSecretManager"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_gcp"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_gitlab", func(r *config.Resource) {
		r.Kind = "SecretSyncGitlab"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_gitlab"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_render", func(r *config.Resource) {
		r.Kind = "SecretSyncRender"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_render"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_spacelift", func(r *config.Resource) {
		r.Kind = "SecretSyncSpacelift"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_spacelift"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_secret_sync_supabase", func(r *config.Resource) {
		r.Kind = "SecretSyncSupabase"
		r.ShortGroup = "secretsync"
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_supabase"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
