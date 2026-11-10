package appconnection

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_app_connection_azure_client_secrets", func(r *config.Resource) {
		r.Kind = "AppConnectionAzureClientSecrets"
		r.ShortGroup = "appconnection" // lowercase not allowed

		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
	})

	p.AddResourceConfigurator("infisical_app_connection_1password", func(r *config.Resource) {
		r.Kind = "AppConnectionOnePassword"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_aws", func(r *config.Resource) {
		r.Kind = "AppConnectionAws"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_azure_app_configuration", func(r *config.Resource) {
		r.Kind = "AppConnectionAzureAppConfiguration"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_azure_devops", func(r *config.Resource) {
		r.Kind = "AppConnectionAzureDevops"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_azure_key_vault", func(r *config.Resource) {
		r.Kind = "AppConnectionAzureKeyVault"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_bitbucket", func(r *config.Resource) {
		r.Kind = "AppConnectionBitbucket"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_chef", func(r *config.Resource) {
		r.Kind = "AppConnectionChef"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_circleci", func(r *config.Resource) {
		r.Kind = "AppConnectionCircleci"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_cloudflare", func(r *config.Resource) {
		r.Kind = "AppConnectionCloudflare"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_databricks", func(r *config.Resource) {
		r.Kind = "AppConnectionDatabricks"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_datadog", func(r *config.Resource) {
		r.Kind = "AppConnectionDatadog"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_digicert", func(r *config.Resource) {
		r.Kind = "AppConnectionDigicert"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_f5_big_ip", func(r *config.Resource) {
		r.Kind = "AppConnectionF5BigIp"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_flyio", func(r *config.Resource) {
		r.Kind = "AppConnectionFlyio"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_gcp", func(r *config.Resource) {
		r.Kind = "AppConnectionGcp"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_github", func(r *config.Resource) {
		r.Kind = "AppConnectionGithub"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_gitlab", func(r *config.Resource) {
		r.Kind = "AppConnectionGitlab"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_hashicorp_vault", func(r *config.Resource) {
		r.Kind = "AppConnectionHashicorpVault"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_kemp_loadmaster", func(r *config.Resource) {
		r.Kind = "AppConnectionKempLoadmaster"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_ldap", func(r *config.Resource) {
		r.Kind = "AppConnectionLdap"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_litellm", func(r *config.Resource) {
		r.Kind = "AppConnectionLitellm"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_mssql", func(r *config.Resource) {
		r.Kind = "AppConnectionMssql"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_mysql", func(r *config.Resource) {
		r.Kind = "AppConnectionMysql"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_netscaler", func(r *config.Resource) {
		r.Kind = "AppConnectionNetscaler"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_nutanix_prism_central", func(r *config.Resource) {
		r.Kind = "AppConnectionNutanixPrismCentral"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_oracledb", func(r *config.Resource) {
		r.Kind = "AppConnectionOracledb"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_postgres", func(r *config.Resource) {
		r.Kind = "AppConnectionPostgres"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_render", func(r *config.Resource) {
		r.Kind = "AppConnectionRender"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_spacelift", func(r *config.Resource) {
		r.Kind = "AppConnectionSpacelift"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_ssh", func(r *config.Resource) {
		r.Kind = "AppConnectionSsh"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_supabase", func(r *config.Resource) {
		r.Kind = "AppConnectionSupabase"
		r.ShortGroup = "appconnection"
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})

	p.AddResourceConfigurator("infisical_app_connection_winrm", func(r *config.Resource) {
		r.Kind = "AppConnectionWinrm"
		r.ShortGroup = "appconnection"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
		r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
	})
}
