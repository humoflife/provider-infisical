package certmanager

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the certmanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_cert_manager_application", func(r *config.Resource) {
		r.Kind = "Application"
		r.ShortGroup = "certmanager"
	})

	p.AddResourceConfigurator("infisical_cert_manager_application_approval_policy", func(r *config.Resource) {
		r.Kind = "ApplicationApprovalPolicy"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_application_group", func(r *config.Resource) {
		r.Kind = "ApplicationGroup"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["group_id"] = config.Reference{TerraformName: "infisical_group"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_application_identity", func(r *config.Resource) {
		r.Kind = "ApplicationIdentity"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_application_profile", func(r *config.Resource) {
		r.Kind = "ApplicationProfile"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["profile_id"] = config.Reference{TerraformName: "infisical_cert_manager_certificate_profile"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_application_user", func(r *config.Resource) {
		r.Kind = "ApplicationUser"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_ca_certificate", func(r *config.Resource) {
		r.Kind = "CaCertificate"
		r.ShortGroup = "certmanager"
		r.References["ca_id"] = config.Reference{TerraformName: "infisical_cert_manager_internal_ca"}
		r.References["parent_ca_id"] = config.Reference{TerraformName: "infisical_cert_manager_internal_ca"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_certificate", func(r *config.Resource) {
		r.Kind = "Certificate"
		r.ShortGroup = "certmanager"
		r.References["application_id"] = config.Reference{TerraformName: "infisical_cert_manager_application"}
		r.References["profile_id"] = config.Reference{TerraformName: "infisical_cert_manager_certificate_profile"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_certificate_policy", func(r *config.Resource) {
		r.Kind = "CertificatePolicy"
		r.ShortGroup = "certmanager"
	})

	p.AddResourceConfigurator("infisical_cert_manager_certificate_profile", func(r *config.Resource) {
		r.Kind = "CertificateProfile"
		r.ShortGroup = "certmanager"
		r.References["certificate_policy_id"] = config.Reference{TerraformName: "infisical_cert_manager_certificate_policy"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_external_ca_acme", func(r *config.Resource) {
		r.Kind = "ExternalCaAcme"
		r.ShortGroup = "certmanager"
	})

	p.AddResourceConfigurator("infisical_cert_manager_external_ca_adcs", func(r *config.Resource) {
		r.Kind = "ExternalCaAdcs"
		r.ShortGroup = "certmanager"
	})

	p.AddResourceConfigurator("infisical_cert_manager_external_ca_digicert", func(r *config.Resource) {
		r.Kind = "ExternalCaDigicert"
		r.ShortGroup = "certmanager"
		r.References["app_connection_id"] = config.Reference{TerraformName: "infisical_app_connection_digicert"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_group", func(r *config.Resource) {
		r.Kind = "Group"
		r.ShortGroup = "certmanager"
		r.References["group_id"] = config.Reference{TerraformName: "infisical_group"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_identity", func(r *config.Resource) {
		r.Kind = "Identity"
		r.ShortGroup = "certmanager"
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})

	p.AddResourceConfigurator("infisical_cert_manager_internal_ca", func(r *config.Resource) {
		r.Kind = "InternalCa"
		r.ShortGroup = "certmanager"
	})

	p.AddResourceConfigurator("infisical_cert_manager_user", func(r *config.Resource) {
		r.Kind = "User"
		r.ShortGroup = "certmanager"
	})
}
