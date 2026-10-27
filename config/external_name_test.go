package config

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func TestWithPlaceholderID(t *testing.T) {
	e := withPlaceholderID(config.IdentifierFromProvider)

	cases := map[string]struct {
		externalName string
		wantID       string
	}{
		"NotCreatedYet": {externalName: "", wantID: emptyIDPlaceholder},
		"Created":       {externalName: "426627eb-1d56-4125-a60a-275633db7574", wantID: "426627eb-1d56-4125-a60a-275633db7574"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			id, err := e.GetIDFn(context.Background(), tc.externalName, nil, nil)
			if err != nil {
				t.Fatalf("GetIDFn: unexpected error: %v", err)
			}
			if id != tc.wantID {
				t.Errorf("GetIDFn: want %q, got %q", tc.wantID, id)
			}
		})
	}

	if _, err := e.GetExternalNameFn(map[string]any{"id": emptyIDPlaceholder}); err == nil {
		t.Error("GetExternalNameFn: want an error for the placeholder ID, got none")
	}
	if _, err := e.GetExternalNameFn(map[string]any{"id": ""}); err == nil {
		t.Error("GetExternalNameFn: want an error for an empty ID, got none")
	}
	name, err := e.GetExternalNameFn(map[string]any{"id": "426627eb-1d56-4125-a60a-275633db7574"})
	if err != nil || name != "426627eb-1d56-4125-a60a-275633db7574" {
		t.Errorf("GetExternalNameFn: want the real ID, got %q, %v", name, err)
	}
}

func TestIdentifierFromAttribute(t *testing.T) {
	const integrationID = "426627eb-1d56-4125-a60a-275633db7574"

	cases := map[string]struct {
		placeholderWhenEmpty bool
		externalName         string
		want                 map[string]any
	}{
		"NotCreatedYetWithPlaceholder":    {placeholderWhenEmpty: true, externalName: "", want: map[string]any{"integration_id": emptyIDPlaceholder}},
		"NotCreatedYetWithoutPlaceholder": {placeholderWhenEmpty: false, externalName: "", want: map[string]any{}},
		"Created":                         {placeholderWhenEmpty: true, externalName: integrationID, want: map[string]any{"integration_id": integrationID}},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := identifierFromAttribute("integration_id", tc.placeholderWhenEmpty)
			state := map[string]any{}
			e.SetIdentifierArgumentFn(state, tc.externalName)
			if len(state) != len(tc.want) || state["integration_id"] != tc.want["integration_id"] {
				t.Errorf("SetIdentifierArgumentFn: want %v, got %v", tc.want, state)
			}
		})
	}

	e := identifierFromAttribute("integration_id", true)
	for _, value := range []any{nil, "", emptyIDPlaceholder} {
		if _, err := e.GetExternalNameFn(map[string]any{"integration_id": value}); err == nil {
			t.Errorf("GetExternalNameFn: want an error for %v, got none", value)
		}
	}
	name, err := e.GetExternalNameFn(map[string]any{"integration_id": integrationID})
	if err != nil || name != integrationID {
		t.Errorf("GetExternalNameFn: want the integration ID, got %q, %v", name, err)
	}
}

func TestGroupMachineIdentityExternalName(t *testing.T) {
	name, err := groupMachineIdentityExternalName.GetExternalNameFn(map[string]any{"group_id": "g", "identity_id": "i"})
	if err != nil || name != "g/i" {
		t.Errorf("GetExternalNameFn: want \"g/i\", got %q, %v", name, err)
	}
	if _, err := groupMachineIdentityExternalName.GetExternalNameFn(map[string]any{"group_id": "g"}); err == nil {
		t.Error("GetExternalNameFn: want an error without identity_id, got none")
	}
}
