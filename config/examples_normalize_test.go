package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
)

func TestNormalizeParameters(t *testing.T) {
	schemas := map[string]*schema.Schema{
		"validity": {Type: conversiontfjson.SchemaTypeObject, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"max":  {Type: schema.TypeString},
			"days": {Type: schema.TypeInt},
		}}},
		"subject": {Type: schema.TypeList, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"type": {Type: schema.TypeString},
		}}},
		"ttl":       {Type: schema.TypeInt},
		"enabled":   {Type: schema.TypeBool},
		"name":      {Type: schema.TypeString},
		"reference": {Type: schema.TypeInt},
	}
	params := map[string]any{
		"validity":  []any{map[string]any{"max": "90d", "days": "30", "unknown": "x"}},
		"subject":   []any{map[string]any{"type": "common_name", "unknown": "x"}},
		"ttl":       "${2592000 * 2}",
		"enabled":   "true",
		"name":      "3600",
		"reference": "${infisical_project.example.ttl}",
		"position":  1,
	}
	want := map[string]any{
		"validity":  map[string]any{"max": "90d", "days": float64(30)},
		"subject":   []any{map[string]any{"type": "common_name"}},
		"ttl":       float64(5184000),
		"enabled":   true,
		"name":      "3600",
		"reference": "${infisical_project.example.ttl}",
	}
	normalizeParameters(params, schemas)
	if diff := cmp.Diff(want, params); diff != "" {
		t.Errorf("normalizeParameters: -want, +got:\n%s", diff)
	}
}
