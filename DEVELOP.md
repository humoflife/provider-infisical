# Develop

## Add a new resource from Terraform

The example adds `infisical_secret_sync_aws_parameter_store`.

1. **Find the resource name** in `config/schema.json`.
   If the resource is not there, increase `TERRAFORM_PROVIDER_VERSION` in the
   `Makefile`.
   Every resource of `config/schema.json` must be configured:
   `TestProviders` in `config/provider_test.go` fails for a resource that is
   not.

2. **Add the external name** in `config/external_name.go`:

   ```go
   "infisical_secret_sync_aws_parameter_store": config.IdentifierFromProvider,
   ```

   If the Terraform Read fails when the ID is empty, use
   `withPlaceholderID(config.IdentifierFromProvider)`.

3. **Set the kind, group and references** in `config/<group>/config.go`:

   ```go
   p.AddResourceConfigurator("infisical_secret_sync_aws_parameter_store", func(r *config.Resource) {
       r.Kind = "SecretSyncAwsParameterStore"
       r.ShortGroup = "secretsync"
       r.References["project_id"] = config.Reference{TerraformName: "infisical_project"}
   })
   ```

   For a new group, add a new `config/<group>/` package, and add its
   `Configure` to `config/provider.go`.

4. **Generate the code:**

   ```sh
   make generate
   ```

   This generates a cluster scoped kind in `apis/cluster/` and a namespaced
   kind in `apis/namespaced/`, both with the API version `v1beta1`.

5. **Build and test:**

   ```sh
   make reviewable
   make build
   ```

6. **Add e2e tests** (optional): add a cluster scoped fixture to
   `test/e2e/testdata/` and an entry to `test/e2e/objects_test.go`. The test
   also creates the namespaced kind from the same fixture.

7. **Commit** the config changes and all the generated files.

## Update the Terraform provider

The provider runs the Terraform provider in-process (no-fork): it imports the
Terraform Plugin Framework provider as the Go module
`terraform-provider-infisical`. The provider code of that module is internal,
so `go.mod` replaces it with the fork `github.com/humoflife/terraform-provider-infisical`,
which adds the exported package `xpprovider` on the branch `xpprovider`. The
fork tags are named after the release they are based on, for example
`v0.20.6-xp.1`.

1. Set `TERRAFORM_PROVIDER_VERSION` in the `Makefile`. `make generate` takes
   the schema and the docs from this release.
2. Rebase the `xpprovider` branch of the fork on the same release tag, tag it,
   and point the `replace` in `go.mod` to that tag. Then run `go mod tidy`.
3. Run `make generate`.
4. Look at the changes in `package/crds/`. If a field of an existing kind is
   removed or changes its type, that is a breaking change for users.

## Release

Packages are only published for a release tag. Merges to `main` never publish.

1. Create a release tag (for example `v0.2.0`): push it with
   `git tag v0.2.0 && git push origin v0.2.0`, or create a new GitHub release
   with a new tag.
2. The **Publish** workflow starts for the new tag, builds the tagged commit,
   and pushes it as that version. It cannot be started by hand.

The version is always the tag name. Publish fails, and pushes nothing, when the
tag is not `vMAJOR.MINOR.PATCH`, or when the version is already in the
registry. A published version is never overwritten: to fix a release, publish
a new version. `make publish` is turned off.
