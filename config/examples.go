package config

import (
	"encoding/json"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

// normalizeExamples makes the examples that the scraper took from the
// Terraform docs match the API types, before upjet generates the example
// manifests from them:
//
//   - the scraper writes every HCL block as a list, but a block with a single
//     nesting is an object in the API, so a list with one object is unwrapped;
//   - HCL accepts numbers and booleans in quotes and constant expressions
//     such as "2592000 * 2", the API does not, so they are evaluated;
//   - attributes that the Terraform schema does not have are mistakes in the
//     docs, so they are removed.
//
// The examples of the dependencies of an example are normalized with the
// schema of their own resource.
func normalizeExamples(pc *ujconfig.Provider) error {
	for _, r := range pc.Resources {
		if r.MetaResource == nil {
			continue
		}
		for i := range r.MetaResource.Examples {
			example := &r.MetaResource.Examples[i]
			normalizeParameters(example.Paved.UnstructuredContent(), r.TerraformResource.Schema)
			for dependencyName, manifest := range example.Dependencies {
				terraformName, _, _ := strings.Cut(dependencyName, ".")
				dependency, ok := pc.Resources[terraformName]
				if !ok {
					continue
				}
				var params map[string]any
				if err := json.Unmarshal([]byte(manifest), &params); err != nil {
					return errors.Wrapf(err, "cannot parse the example of %s in the examples of %s", dependencyName, r.Name)
				}
				normalizeParameters(params, dependency.TerraformResource.Schema)
				normalized, err := json.Marshal(params)
				if err != nil {
					return errors.Wrapf(err, "cannot write the example of %s in the examples of %s", dependencyName, r.Name)
				}
				example.Dependencies[dependencyName] = string(normalized)
			}
		}
	}
	return nil
}

// normalizeParameters normalizes the parameters of an example in place.
func normalizeParameters(params map[string]any, schemas map[string]*schema.Schema) {
	for key, value := range params {
		s, ok := schemas[key]
		if !ok {
			delete(params, key)
			continue
		}
		params[key] = normalizeValue(value, s)
	}
}

func normalizeValue(value any, s *schema.Schema) any {
	switch s.Type { //nolint:exhaustive // the other types need no normalization
	case conversiontfjson.SchemaTypeObject:
		if list, ok := value.([]any); ok && len(list) == 1 {
			value = list[0]
		}
		normalizeNested(value, s)
	case schema.TypeList, schema.TypeSet:
		if list, ok := value.([]any); ok {
			for _, element := range list {
				normalizeNested(element, s)
			}
		}
	case schema.TypeInt, schema.TypeFloat:
		if v, ok := evaluateConstant(value, cty.Number); ok {
			number, _ := v.AsBigFloat().Float64()
			return number
		}
	case schema.TypeBool:
		if v, ok := evaluateConstant(value, cty.Bool); ok {
			return v.True()
		}
	}
	return value
}

// evaluateConstant evaluates a string value of an example as an HCL
// expression without variables, for example "3600", "true" or
// "${2592000 * 2}", and converts the result to the wanted type. It returns
// false for values that are not strings or not constant expressions.
func evaluateConstant(value any, wantType cty.Type) (cty.Value, bool) {
	text, ok := value.(string)
	if !ok {
		return cty.NilVal, false
	}
	if inner, isTemplate := strings.CutPrefix(text, "${"); isTemplate {
		text = strings.TrimSuffix(inner, "}")
	}
	expression, diagnostics := hclsyntax.ParseExpression([]byte(text), "example", hcl.InitialPos)
	if diagnostics.HasErrors() {
		return cty.NilVal, false
	}
	v, diagnostics := expression.Value(nil)
	if diagnostics.HasErrors() {
		return cty.NilVal, false
	}
	v, err := convert.Convert(v, wantType)
	if err != nil || v.IsNull() || !v.IsKnown() {
		return cty.NilVal, false
	}
	return v, true
}

// normalizeNested normalizes an object of a nested block or attribute.
func normalizeNested(value any, s *schema.Schema) {
	object, ok := value.(map[string]any)
	if !ok {
		return
	}
	if resource, ok := s.Elem.(*schema.Resource); ok {
		normalizeParameters(object, resource.Schema)
	}
}
