/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/infisical/provider-infisical/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/infisical/provider-infisical/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal infisical credentials as JSON"
)

// TerraformSetupBuilder builds a terraform.SetupFn function which returns the
// configuration of the given Terraform Plugin Framework provider. The provider
// runs in-process and upjet configures it with ps.Configuration.
func TerraformSetupBuilder(frameworkProvider fwprovider.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			FrameworkProvider: frameworkProvider,
			Configuration:     map[string]any{},
		}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		if pcSpec.Host != "" {
			ps.Configuration["host"] = pcSpec.Host
		}

		creds := map[string]any{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		authField, ok := creds["auth"]
		if !ok || authField == nil {
			return ps, errors.Wrap(fmt.Errorf("auth is not set"), errUnmarshalCredentials)
		}

		authMap, ok := authField.(map[string]any)
		if !ok {
			return ps, errors.Wrap(fmt.Errorf("auth is not a valid object"), errUnmarshalCredentials)
		}

		if universal, ok := authMap["universal"]; ok && universal != nil {
			universalMap, ok := universal.(map[string]any)
			if !ok {
				return ps, errors.Wrap(fmt.Errorf("universal auth is not a valid object"), errUnmarshalCredentials)
			}

			clientID, hasClientID := universalMap["client_id"]
			clientSecret, hasClientSecret := universalMap["client_secret"]

			if !hasClientID || !hasClientSecret {
				return ps, errors.Wrap(fmt.Errorf("missing client_id or client_secret in universal auth"), errUnmarshalCredentials)
			}

			ps.Configuration["auth"] = map[string]any{
				"universal": map[string]any{
					"client_id":     clientID,
					"client_secret": clientSecret,
				},
			}
		}

		if token, ok := authMap["token"]; ok && token != nil {
			ps.Configuration["auth"] = map[string]any{
				"token": token,
			}
		}

		if kubernetes, ok := authMap["kubernetes"]; ok && kubernetes != nil {
			kubernetesMap, ok := kubernetes.(map[string]any)
			if !ok {
				return ps, errors.Wrap(fmt.Errorf("kubernetes auth is not a valid object"), errUnmarshalCredentials)
			}

			identityID, hasIdentityID := kubernetesMap["identity_id"]
			if !hasIdentityID {
				return ps, errors.Wrap(fmt.Errorf("missing identity_id in kubernetes auth"), errUnmarshalCredentials)
			}

			kubeMap := map[string]any{
				"identity_id": identityID,
			}

			serviceAccountTokenPath, hasServiceAccountTokenPath := kubernetesMap["service_account_token_path"]
			if hasServiceAccountTokenPath {
				kubeMap["service_account_token_path"] = serviceAccountTokenPath
			}

			ps.Configuration["auth"] = map[string]any{
				"kubernetes": kubeMap,
			}
		}

		return ps, nil
	}
}

// resolveProviderConfig returns the spec of the ProviderConfig of the managed
// resource, and tracks its usage. Cluster scoped managed resources reference
// a cluster scoped ProviderConfig, namespaced managed resources reference a
// ProviderConfig in their namespace or a ClusterProviderConfig.
func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, crClient client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(crClient, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.Errorf("%s is not an Object", configRef.Kind)
	}

	// a ProviderConfig is in the namespace of the managed resource, a
	// ClusterProviderConfig is cluster scoped
	key := types.NamespacedName{Name: configRef.Name}
	if configRef.Kind == namespacedv1beta1.ProviderConfigKind {
		key.Namespace = mg.GetNamespace()
	}
	if err := crClient.Get(ctx, key, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		// a namespaced ProviderConfig can only use a Secret in its own namespace
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}

// toSharedPCSpec converts the spec of a cluster scoped ProviderConfig to the
// spec of the namespaced one. Both have the same fields.
func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var spec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &spec)
	return &spec, err
}
