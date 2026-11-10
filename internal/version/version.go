// Package version contains values that the build sets with -ldflags.
package version

var (
	// Version of the provider.
	Version string

	// TerraformProviderVersion is the version of the Terraform provider that
	// runs in-process. The build sets it from TERRAFORM_PROVIDER_VERSION in
	// the Makefile.
	TerraformProviderVersion string
)
