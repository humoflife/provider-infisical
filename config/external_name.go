/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/pkg/errors"
)

// emptyIDPlaceholder is the Terraform ID of a resource that is not created yet
const emptyIDPlaceholder = "00000000-0000-0000-0000-000000000000"

// returns a copy of the external name configuration that uses emptyIDPlaceholder as the Terraform ID while the external name is empty
func withPlaceholderID(e config.ExternalName) config.ExternalName {
	getID := e.GetIDFn
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		if externalName == "" {
			return emptyIDPlaceholder, nil
		}
		return getID(ctx, externalName, parameters, providerConfig)
	}
	getExternalName := e.GetExternalNameFn
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		name, err := getExternalName(tfstate)
		if err == nil && name == emptyIDPlaceholder {
			return "", errors.New("cannot find id in tfstate")
		}
		return name, err
	}
	return e
}

// identifierFromAttribute returns an external name configuration for a
// resource without an id attribute: the external name is the value of the
// given attribute, which the Terraform Read uses to find the resource. With
// placeholderWhenEmpty, the Read gets emptyIDPlaceholder while the external
// name is empty, so that it looks up a resource that does not exist instead of
// one with an empty ID.
func identifierFromAttribute(attribute string, placeholderWhenEmpty bool) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		value, ok := tfstate[attribute].(string)
		if !ok || value == "" || value == emptyIDPlaceholder {
			return "", errors.Errorf("cannot find %s in tfstate", attribute)
		}
		return value, nil
	}
	e.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
		switch {
		case externalName != "":
			base[attribute] = externalName
		case placeholderWhenEmpty:
			base[attribute] = emptyIDPlaceholder
		}
	}
	return e
}

// groupMachineIdentityExternalName is the external name configuration of
// infisical_group_machine_identity, which has no id attribute. Its Terraform
// Read finds the membership with the group_id and identity_id parameters, so
// the external name is only informational: "<group_id>/<identity_id>".
var groupMachineIdentityExternalName = func() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		groupID, groupOK := tfstate["group_id"].(string)
		identityID, identityOK := tfstate["identity_id"].(string)
		if !groupOK || !identityOK || groupID == "" || identityID == "" {
			return "", errors.New("cannot find group_id and identity_id in tfstate")
		}
		return groupID + "/" + identityID, nil
	}
	return e
}()

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	"infisical_secret_rotation_azure_client_secret": withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_azure_client_secrets": withPlaceholderID(config.IdentifierFromProvider),

	"infisical_project":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity":                 config.IdentifierFromProvider,
	"infisical_group":                    config.IdentifierFromProvider,
	"infisical_project_environment":      config.IdentifierFromProvider,
	"infisical_identity_universal_auth":  config.IdentifierFromProvider,
	"infisical_identity_kubernetes_auth": config.IdentifierFromProvider,
	"infisical_project_role":             config.IdentifierFromProvider,
	"infisical_project_template":         config.IdentifierFromProvider,
	"infisical_secret_sync_github":       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_access_approval_policy":   withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_approval_policy":   withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_folder":            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_project_identity": func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			membershipID, ok := tfstate["membership_id"]
			if !ok {
				return "", errors.Errorf("no membership_id attribute found in tfstate")
			}

			if membershipID == nil {
				return "", nil
			}

			membershipIDStr, ok := membershipID.(string)
			if !ok {
				return "", errors.Errorf("membership_id attribute is not a string")
			}
			return membershipIDStr, nil
		}
		return e
	}(),
	"infisical_project_user": func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			membershipID, ok := tfstate["membership_id"]
			if !ok {
				return "", errors.Errorf("no membership_id attribute found in tfstate")
			}

			if membershipID == nil {
				return "", nil
			}

			membershipIDStr, ok := membershipID.(string)
			if !ok {
				return "", errors.Errorf("membership_id attribute is not a string")
			}
			return membershipIDStr, nil
		}
		return e
	}(),
	"infisical_project_group": func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			membershipID, ok := tfstate["membership_id"]
			if !ok {
				return "", errors.Errorf("no membership_id attribute found in tfstate")
			}

			if membershipID == nil {
				return "", nil
			}

			membershipIDStr, ok := membershipID.(string)
			if !ok {
				return "", errors.Errorf("membership_id attribute is not a string")
			}
			return membershipIDStr, nil
		}
		return e
	}(),

	"infisical_secret": withPlaceholderID(func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["id"]; ok && id != nil {
				idStr, ok := id.(string)
				if ok && idStr != "" {
					return idStr, nil
				}
			}
			return "", nil
		}
		return e
	}()),

	// The resources below were added with the no-fork provider. They have
	// not been tested against Infisical yet.
	"infisical_alert":                                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_1password":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_aws":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_azure_app_configuration":         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_azure_devops":                    withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_azure_key_vault":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_bitbucket":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_chef":                            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_circleci":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_cloudflare":                      withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_databricks":                      withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_datadog":                         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_digicert":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_f5_big_ip":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_flyio":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_gcp":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_github":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_gitlab":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_hashicorp_vault":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_kemp_loadmaster":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_ldap":                            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_litellm":                         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_mssql":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_mysql":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_netscaler":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_nutanix_prism_central":           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_oracledb":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_postgres":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_render":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_spacelift":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_ssh":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_supabase":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_app_connection_winrm":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application_approval_policy":       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application_group":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application_identity":              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application_profile":               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_application_user":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_ca_certificate":                    withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_certificate":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_certificate_policy":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_certificate_profile":               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_external_ca_acme":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_external_ca_adcs":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_external_ca_digicert":              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_group":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_identity":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_internal_ca":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_cert_manager_user":                              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_aws_certificate_manager":       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_aws_elastic_load_balancer":     withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_aws_secrets_manager":           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_azure_key_vault":               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_certificate":                   withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_chef":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_cloudflare_custom_certificate": withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_f5_big_ip":                     withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_gcp_certificate_manager":       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_kemp_loadmaster":               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_linux_server":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_netscaler":                     withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_nutanix_prism_central":         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_certificate_sync_windows_server":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_dynamic_secret_aws_iam":                         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_dynamic_secret_kubernetes":                      withPlaceholderID(config.IdentifierFromProvider),
	"infisical_dynamic_secret_mongo_atlas":                     withPlaceholderID(config.IdentifierFromProvider),
	"infisical_dynamic_secret_mongo_db":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_dynamic_secret_sql_database":                    withPlaceholderID(config.IdentifierFromProvider),
	"infisical_external_kms_aws":                               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_gateway":                                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_gateway_enrollment_token":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_group_machine_identity":                         groupMachineIdentityExternalName,
	"infisical_identity_aws_auth":                              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_azure_auth":                            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_gcp_auth":                              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_jwt_auth":                              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_kubernetes_auth_template":              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_oidc_auth":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_oidc_auth_template":                    withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_tls_cert_auth":                         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_token_auth":                            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_token_auth_token":                      withPlaceholderID(config.IdentifierFromProvider),
	"infisical_identity_universal_auth_client_secret":          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_integration_aws_parameter_store":                identifierFromAttribute("integration_id", true),
	"infisical_integration_aws_secrets_manager":                identifierFromAttribute("integration_id", true),
	"infisical_integration_circleci":                           identifierFromAttribute("integration_id", true),
	"infisical_integration_databricks":                         identifierFromAttribute("integration_id", true),
	"infisical_integration_gcp_secret_manager":                 identifierFromAttribute("integration_id", true),
	"infisical_kms_key":                                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_org_role":                                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_project_identity_specific_privilege":            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_project_scoped_identity":                        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_import":                                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_auth0_client_secret":            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_aws_iam_user_secret":            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_datadog_api_key":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_ldap_password":                  withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_litellm_api_key":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_mssql_credentials":              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_mysql_credentials":              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_oracledb_credentials":           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_rotation_postgres_credentials":           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_1password":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_aws_parameter_store":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_aws_secrets_manager":                withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_azure_app_configuration":            withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_azure_devops":                       withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_azure_key_vault":                    withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_bitbucket":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_circleci":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_cloudflare_pages":                   withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_cloudflare_workers":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_databricks":                         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_flyio":                              withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_gcp_secret_manager":                 withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_gitlab":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_render":                             withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_spacelift":                          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_sync_supabase":                           withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_tag":                                     withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_validation_rule_dynamic_secrets":         withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_validation_rule_secret_rotations":        withPlaceholderID(config.IdentifierFromProvider),
	"infisical_secret_validation_rule_static_secrets":          withPlaceholderID(config.IdentifierFromProvider),
	"infisical_sub_organization":                               withPlaceholderID(config.IdentifierFromProvider),
	"infisical_sub_organization_group":                         identifierFromAttribute("membership_id", false),
	"infisical_webhook":                                        withPlaceholderID(config.IdentifierFromProvider),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
