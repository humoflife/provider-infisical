package config

import (
	"encoding/json"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// TestProviders checks that every Terraform resource of the schema is
// configured in both scopes, with a kind and a group, and is reconciled by the
// in-process Terraform Plugin Framework provider. NewProvider panics when a
// configured resource is not in the Go provider.
func TestProviders(t *testing.T) {
	var schema struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]json.RawMessage `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal([]byte(providerSchema), &schema); err != nil {
		t.Fatal(err)
	}
	var terraformResources []string
	for _, ps := range schema.ProviderSchemas {
		for name := range ps.ResourceSchemas {
			terraformResources = append(terraformResources, name)
		}
	}

	for scope, provider := range map[string]*ujconfig.Provider{"cluster": GetProvider(), "namespaced": GetProviderNamespaced()} {
		t.Run(scope, func(t *testing.T) {
			if len(provider.Resources) != len(terraformResources) {
				t.Errorf("want %d resources, got %d", len(terraformResources), len(provider.Resources))
			}
			kinds := map[string]string{}
			for _, name := range terraformResources {
				r, ok := provider.Resources[name]
				if !ok {
					t.Errorf("%s is not configured", name)
					continue
				}
				if !r.ShouldUseTerraformPluginFrameworkClient() {
					t.Errorf("%s is not reconciled with the Terraform Plugin Framework provider", name)
				}
				if r.Version != "v1beta1" {
					t.Errorf("%s: want version v1beta1, got %s", name, r.Version)
				}
				key := r.ShortGroup + "/" + r.Kind
				if other, ok := kinds[key]; ok {
					t.Errorf("%s and %s have the same group and kind %s", name, other, key)
				}
				kinds[key] = name
			}
		})
	}
}
