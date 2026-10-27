package identity

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity", func(r *config.Resource) {
		r.Kind = "Identity"
		r.ShortGroup = "identity"
	})

	p.AddResourceConfigurator("infisical_identity_universal_auth", func(r *config.Resource) {
		r.Kind = "UniversalAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
	})

	p.AddResourceConfigurator("infisical_identity_kubernetes_auth", func(r *config.Resource) {
		r.Kind = "KubernetesAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
	})

	p.AddResourceConfigurator("infisical_identity_aws_auth", func(r *config.Resource) {
		r.Kind = "AwsAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_azure_auth", func(r *config.Resource) {
		r.Kind = "AzureAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_gcp_auth", func(r *config.Resource) {
		r.Kind = "GcpAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_jwt_auth", func(r *config.Resource) {
		r.Kind = "JwtAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_kubernetes_auth_template", func(r *config.Resource) {
		r.Kind = "KubernetesAuthTemplate"
		r.ShortGroup = "identity"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
	})

	p.AddResourceConfigurator("infisical_identity_oidc_auth", func(r *config.Resource) {
		r.Kind = "OidcAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
		r.References["template_id"] = config.Reference{TerraformName: "infisical_identity_oidc_auth_template"}
	})

	p.AddResourceConfigurator("infisical_identity_oidc_auth_template", func(r *config.Resource) {
		r.Kind = "OidcAuthTemplate"
		r.ShortGroup = "identity"
	})

	p.AddResourceConfigurator("infisical_identity_tls_cert_auth", func(r *config.Resource) {
		r.Kind = "TlsCertAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_token_auth", func(r *config.Resource) {
		r.Kind = "TokenAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_token_auth_token", func(r *config.Resource) {
		r.Kind = "TokenAuthToken"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_identity_universal_auth_client_secret", func(r *config.Resource) {
		r.Kind = "UniversalAuthClientSecret"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})
}
