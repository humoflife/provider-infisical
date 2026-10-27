package config

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	ujresource "github.com/crossplane/upjet/v2/pkg/resource"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	apisCluster "github.com/infisical/provider-infisical/apis/cluster"
	apisNamespaced "github.com/infisical/provider-infisical/apis/namespaced"
)

// TestExamplesMatchTerraformSchema decodes every generated example into its
// API type and converts its parameters to a value of the schema of the
// in-process Terraform resource, the same way upjet does before it calls the
// Terraform provider. It finds API types that do not match the Terraform
// schema without calling Infisical.
func TestExamplesMatchTerraformSchema(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := apisCluster.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	if err := apisNamespaced.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	decoder := serializer.NewCodecFactory(scheme).UniversalDeserializer()

	for scope, provider := range map[string]*ujconfig.Provider{"cluster": GetProvider(), "namespaced": GetProviderNamespaced()} {
		root := filepath.Join("..", "examples-generated", scope)
		tested := map[string]bool{}
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".yaml") {
				return err
			}
			relativePath, _ := filepath.Rel(root, path)
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			for i, document := range strings.Split(string(raw), "\n---\n") {
				object, _, err := decoder.Decode([]byte(document), nil, nil)
				if err != nil {
					t.Errorf("%s/%s, document %d: cannot decode: %v", scope, relativePath, i, err)
					continue
				}
				tr, ok := object.(ujresource.Terraformed)
				if !ok {
					continue
				}
				r, ok := provider.Resources[tr.GetTerraformResourceType()]
				if !ok {
					t.Errorf("%s/%s: %s is not configured", scope, relativePath, tr.GetTerraformResourceType())
					continue
				}
				if err := checkParameters(tr, r); err != nil {
					t.Errorf("%s/%s, document %d (%s): %v", scope, relativePath, i, r.Name, err)
				}
				tested[r.Name] = true
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s: checked the examples of %d of %d resources", scope, len(tested), len(provider.Resources))
	}
}

// checkParameters converts the parameters of the managed resource to a value
// of the Terraform resource schema, like upjet's Terraform Plugin Framework
// connector does for a resource that is not created yet.
func checkParameters(tr ujresource.Terraformed, r *ujconfig.Resource) error {
	params, err := tr.GetMergedParameters(false)
	if err != nil {
		return err
	}
	return checkTerraformParameters(params, r)
}

// checkTerraformParameters converts the parameters to a value of the
// Terraform resource schema.
func checkTerraformParameters(params map[string]any, r *ujconfig.Resource) error {
	ctx := context.Background()
	var err error
	if params, err = r.ApplyTFConversions(params, ujconfig.ToTerraform); err != nil {
		return err
	}
	r.ExternalName.SetIdentifierArgumentFn(params, "")

	schemaResponse := &fwresource.SchemaResponse{}
	r.TerraformPluginFrameworkResource.Schema(ctx, fwresource.SchemaRequest{}, schemaResponse)
	if schemaResponse.Diagnostics.HasError() {
		return errorsFromDiagnostics(schemaResponse)
	}
	if _, ok := schemaResponse.Schema.Attributes["id"]; ok {
		if params["id"], err = r.ExternalName.GetIDFn(ctx, "", params, nil); err != nil {
			return err
		}
	}
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	_, err = tftypes.ValueFromJSONWithOpts(paramsJSON, schemaResponse.Schema.Type().TerraformType(ctx), tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true})
	return err
}

func errorsFromDiagnostics(response *fwresource.SchemaResponse) error {
	var messages []string
	for _, d := range response.Diagnostics.Errors() {
		messages = append(messages, d.Summary()+": "+d.Detail())
	}
	return errors.New(strings.Join(messages, "; "))
}
