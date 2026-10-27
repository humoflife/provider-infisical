/*
Copyright 2026 Infisical Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/google/go-cmp/cmp"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"terraform-provider-infisical/xpprovider"

	clusterproject "github.com/infisical/provider-infisical/apis/cluster/project/v1beta1"
	clusterv1beta1 "github.com/infisical/provider-infisical/apis/cluster/v1beta1"
	namespacedproject "github.com/infisical/provider-infisical/apis/namespaced/project/v1beta1"
	namespacedv1beta1 "github.com/infisical/provider-infisical/apis/namespaced/v1beta1"
)

const (
	testHost         = "https://infisical.example.com"
	testSecretName   = "infisical-creds"
	testSecretKey    = "credentials"
	testMRNamespace  = "team-a"
	testClientID     = "client-id"
	testClientSecret = "client-secret"
)

var universalAuthCredentials = `{"auth":{"universal":{"client_id":"` + testClientID + `","client_secret":"` + testClientSecret + `"}}}`

var wantUniversalAuthConfiguration = map[string]any{
	"host": testHost,
	"auth": map[string]any{
		"universal": map[string]any{
			"client_id":     testClientID,
			"client_secret": testClientSecret,
		},
	},
}

func newScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	scheme := runtime.NewScheme()
	for _, addToScheme := range []func(*runtime.Scheme) error{
		corev1.AddToScheme,
		clusterv1beta1.SchemeBuilder.AddToScheme,
		namespacedv1beta1.SchemeBuilder.AddToScheme,
		clusterproject.SchemeBuilder.AddToScheme,
		namespacedproject.SchemeBuilder.AddToScheme,
	} {
		if err := addToScheme(scheme); err != nil {
			t.Fatal(err)
		}
	}
	return scheme
}

// newRESTMapper returns a REST mapper that knows the scope of the kinds. The
// default REST mapper of the fake client treats every kind as namespaced.
func newRESTMapper(scheme *runtime.Scheme) meta.RESTMapper {
	clusterScopedKinds := map[schema.GroupVersionKind]bool{
		clusterv1beta1.ProviderConfigGroupVersionKind:           true,
		clusterv1beta1.ProviderConfigUsageGroupVersionKind:      true,
		namespacedv1beta1.ClusterProviderConfigGroupVersionKind: true,
		clusterproject.Project_GroupVersionKind:                 true,
	}
	restMapper := meta.NewDefaultRESTMapper(nil)
	for gvk := range scheme.AllKnownTypes() {
		scope := meta.RESTScopeNamespace
		if clusterScopedKinds[gvk] {
			scope = meta.RESTScopeRoot
		}
		restMapper.Add(gvk, scope)
	}
	return restMapper
}

func credentialsSecret(namespace string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: testSecretName, Namespace: namespace},
		Data:       map[string][]byte{testSecretKey: []byte(universalAuthCredentials)},
	}
}

func secretCredentials(secretNamespace string) namespacedv1beta1.ProviderCredentials {
	return namespacedv1beta1.ProviderCredentials{
		Source: xpv2.CredentialsSourceSecret,
		CommonCredentialSelectors: xpv2.CommonCredentialSelectors{
			SecretRef: &xpv2.SecretKeySelector{
				SecretReference: xpv2.SecretReference{Name: testSecretName, Namespace: secretNamespace},
				Key:             testSecretKey,
			},
		},
	}
}

func TestTerraformSetupBuilder(t *testing.T) {
	frameworkProvider := xpprovider.New("test")

	cases := map[string]struct {
		reason  string
		objects []client.Object
		mg      resource.Managed
	}{
		"ClusterScopedManagedResource": {
			reason: "A cluster scoped managed resource uses the cluster scoped ProviderConfig.",
			objects: []client.Object{
				credentialsSecret("crossplane-system"),
				&clusterv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
					Spec: clusterv1beta1.ProviderConfigSpec{
						Host: testHost,
						Credentials: clusterv1beta1.ProviderCredentials{
							Source:                    xpv2.CredentialsSourceSecret,
							CommonCredentialSelectors: secretCredentials("crossplane-system").CommonCredentialSelectors,
						},
					},
				},
			},
			mg: &clusterproject.Project{
				ObjectMeta: metav1.ObjectMeta{Name: "project", UID: "uid-1"},
				Spec: clusterproject.ProjectSpec{
					ClusterManagedResourceSpec: xpv2.ClusterManagedResourceSpec{ProviderConfigReference: &xpv2.Reference{Name: "default"}},
				},
			},
		},
		"NamespacedManagedResourceWithProviderConfig": {
			reason: "A namespaced ProviderConfig reads its Secret from the namespace of the managed resource, whatever namespace it names.",
			objects: []client.Object{
				credentialsSecret(testMRNamespace),
				&namespacedv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: testMRNamespace},
					Spec: namespacedv1beta1.ProviderConfigSpec{
						Host:        testHost,
						Credentials: secretCredentials("some-other-namespace"),
					},
				},
			},
			mg: &namespacedproject.Project{
				ObjectMeta: metav1.ObjectMeta{Name: "project", Namespace: testMRNamespace, UID: "uid-2"},
				Spec: namespacedproject.ProjectSpec{
					ManagedResourceSpec: xpv2.ManagedResourceSpec{
						ProviderConfigReference: &xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: "default"},
					},
				},
			},
		},
		"NamespacedManagedResourceWithClusterProviderConfig": {
			reason: "A ClusterProviderConfig reads its Secret from the namespace it names.",
			objects: []client.Object{
				credentialsSecret("crossplane-system"),
				&namespacedv1beta1.ClusterProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
					Spec: namespacedv1beta1.ProviderConfigSpec{
						Host:        testHost,
						Credentials: secretCredentials("crossplane-system"),
					},
				},
			},
			mg: &namespacedproject.Project{
				ObjectMeta: metav1.ObjectMeta{Name: "project", Namespace: testMRNamespace, UID: "uid-3"},
				Spec: namespacedproject.ProjectSpec{
					ManagedResourceSpec: xpv2.ManagedResourceSpec{
						ProviderConfigReference: &xpv2.ProviderConfigReference{Kind: "ClusterProviderConfig", Name: "default"},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			scheme := newScheme(t)
			kubeClient := fake.NewClientBuilder().WithScheme(scheme).WithRESTMapper(newRESTMapper(scheme)).WithObjects(tc.objects...).Build()

			setup, err := TerraformSetupBuilder(frameworkProvider)(context.Background(), kubeClient, tc.mg)
			if err != nil {
				t.Fatalf("\n%s\nTerraformSetupBuilder(...): unexpected error: %v", tc.reason, err)
			}
			if setup.FrameworkProvider != frameworkProvider {
				t.Errorf("\n%s\nTerraformSetupBuilder(...): the setup does not carry the framework provider", tc.reason)
			}
			if diff := cmp.Diff(wantUniversalAuthConfiguration, map[string]any(setup.Configuration)); diff != "" {
				t.Errorf("\n%s\nTerraformSetupBuilder(...): -want configuration, +got configuration:\n%s", tc.reason, diff)
			}
			assertFitsProviderSchema(t, frameworkProvider, setup.Configuration)
		})
	}
}

// assertFitsProviderSchema checks that upjet can convert the configuration to
// a value of the provider schema, the way it does before it configures the
// in-process provider.
func assertFitsProviderSchema(t *testing.T, frameworkProvider fwprovider.Provider, configuration map[string]any) {
	t.Helper()
	ctx := context.Background()
	schemaResponse := &fwprovider.SchemaResponse{}
	frameworkProvider.Schema(ctx, fwprovider.SchemaRequest{}, schemaResponse)
	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("cannot get the provider schema: %v", schemaResponse.Diagnostics)
	}
	configurationJSON, err := json.Marshal(configuration)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tftypes.ValueFromJSONWithOpts(configurationJSON, schemaResponse.Schema.Type().TerraformType(ctx), tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true}); err != nil {
		t.Errorf("the configuration does not fit the provider schema: %v", err)
	}
}
