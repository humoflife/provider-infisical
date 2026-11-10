//go:build e2e

/*
Copyright 2026 Infisical Inc.
*/

package e2e

import (
	"context"
	"testing"
)

// TestLifecycle creates every kind in Infisical, both as a cluster scoped and
// as a namespaced managed resource, and checks the full lifecycle:
//
//   - create: the object becomes Ready and gets an external name;
//   - update: a change of the object reaches Infisical;
//   - delete: the object and the Infisical resource are deleted.
func TestLifecycle(t *testing.T) {
	for _, sc := range scopes {
		t.Run(string(sc), func(t *testing.T) {
			t.Parallel()
			s := newSet(sc, "e2e-life-"+string(sc)[:4], "default")
			orphan := func(o object) bool { return o.notReady != "" }

			if t.Run("create", func(t *testing.T) { s.createAll(t) }) {
				t.Run("reconcile", func(t *testing.T) {
					for _, o := range objects {
						t.Run(o.testName(), func(t *testing.T) {
							t.Parallel()
							s.waitReconciled(t, o)
							checkExternalName(t, s, o)
						})
					}
				})
				t.Run("update", func(t *testing.T) {
					for _, o := range objects {
						t.Run(o.testName(), func(t *testing.T) {
							t.Parallel()
							if o.notReady != "" {
								t.Skip(o.notReady)
							}
							s.update(t, o)
						})
					}
				})
			}
			t.Run("delete", func(t *testing.T) { s.deleteAll(t, orphan) })
		})
	}
}

// checkExternalName checks that a ready object has an external name, which is
// the ID of the resource in Infisical.
func checkExternalName(t *testing.T, s *set, o object) {
	t.Helper()
	if o.notReady != "" {
		return
	}
	u, err := s.get(context.Background(), o)
	if err != nil {
		t.Fatal(err)
	}
	if u.GetAnnotations()["crossplane.io/external-name"] == "" {
		t.Error("the object is ready but has no external name")
	}
}
