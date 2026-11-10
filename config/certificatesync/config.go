package certificatesync

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the certificatesync group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_certificate_sync_aws_certificate_manager", func(r *config.Resource) {
		r.Kind = "AwsCertificateManager"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_aws_elastic_load_balancer", func(r *config.Resource) {
		r.Kind = "AwsElasticLoadBalancer"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_aws_secrets_manager", func(r *config.Resource) {
		r.Kind = "AwsSecretsManager"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_aws"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_azure_key_vault", func(r *config.Resource) {
		r.Kind = "AzureKeyVault"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_azure_key_vault"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_certificate", func(r *config.Resource) {
		r.Kind = "Certificate"
		r.ShortGroup = "certificatesync"
		r.References["certificate_id"] = config.Reference{TerraformName: "infisical_cert_manager_certificate"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_chef", func(r *config.Resource) {
		r.Kind = "Chef"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_chef"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_cloudflare_custom_certificate", func(r *config.Resource) {
		r.Kind = "CloudflareCustomCertificate"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_cloudflare"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_f5_big_ip", func(r *config.Resource) {
		r.Kind = "F5BigIp"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_f5_big_ip"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_gcp_certificate_manager", func(r *config.Resource) {
		r.Kind = "GcpCertificateManager"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_gcp"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_kemp_loadmaster", func(r *config.Resource) {
		r.Kind = "KempLoadmaster"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_kemp_loadmaster"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_linux_server", func(r *config.Resource) {
		r.Kind = "LinuxServer"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_netscaler", func(r *config.Resource) {
		r.Kind = "Netscaler"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_netscaler"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_nutanix_prism_central", func(r *config.Resource) {
		r.Kind = "NutanixPrismCentral"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["connection_id"] = config.Reference{TerraformName: "infisical_app_connection_nutanix_prism_central"}
	})

	p.AddResourceConfigurator("infisical_certificate_sync_windows_server", func(r *config.Resource) {
		r.Kind = "WindowsServer"
		r.ShortGroup = "certificatesync"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
	})
}
