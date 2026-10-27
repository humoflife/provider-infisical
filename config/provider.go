/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"terraform-provider-infisical/xpprovider"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/infisical/provider-infisical/config/appconnection"
	"github.com/infisical/provider-infisical/config/certificatesync"
	"github.com/infisical/provider-infisical/config/certmanager"
	"github.com/infisical/provider-infisical/config/dynamicsecret"
	"github.com/infisical/provider-infisical/config/gateway"
	"github.com/infisical/provider-infisical/config/group"
	"github.com/infisical/provider-infisical/config/identity"
	"github.com/infisical/provider-infisical/config/integration"
	"github.com/infisical/provider-infisical/config/kms"
	"github.com/infisical/provider-infisical/config/organization"
	"github.com/infisical/provider-infisical/config/project"
	"github.com/infisical/provider-infisical/config/secret"
	"github.com/infisical/provider-infisical/config/secretrotation"
	"github.com/infisical/provider-infisical/config/secretsync"
	"github.com/infisical/provider-infisical/internal/version"
)

const (
	resourcePrefix = "infisical"
	modulePath     = "github.com/infisical/provider-infisical"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the configuration of the cluster scoped managed
// resources.
func GetProvider() *ujconfig.Provider {
	return newProvider("infisical.upbound.io")
}

// GetProviderNamespaced returns the configuration of the namespaced managed
// resources.
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("infisical.m.upbound.io",
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "default",
		}))
}

// newProvider returns the configuration of the managed resources in the given
// root group. All resources are reconciled with the Terraform Plugin Framework
// provider in-process (no-fork), so the provider does not run the Terraform CLI.
func newProvider(rootGroup string, opts ...ujconfig.ProviderOption) *ujconfig.Provider {
	opts = append([]ujconfig.ProviderOption{
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkProvider(xpprovider.New(version.TerraformProviderVersion)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			apiVersion,
		),
	}, opts...)
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata), opts...)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		identity.Configure,
		group.Configure,
		secretsync.Configure,
		secret.Configure,
		secretrotation.Configure,
		appconnection.Configure,
		certificatesync.Configure,
		certmanager.Configure,
		dynamicsecret.Configure,
		gateway.Configure,
		integration.Configure,
		kms.Configure,
		organization.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	if err := normalizeExamples(pc); err != nil {
		panic(err)
	}
	return pc
}

// apiVersion sets the API version of every resource. All resources have the
// single version v1beta1, in both scopes.
func apiVersion(r *ujconfig.Resource) {
	r.Version = "v1beta1"
}
