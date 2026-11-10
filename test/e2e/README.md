# End-to-end tests

The e2e tests install the locally built provider into a kind cluster and run
the full lifecycle (create, update, delete) of every kind against a real
Infisical instance, both as cluster scoped and as namespaced managed resources
(`TestLifecycle`). They run with Crossplane 1.16.5 and 2.4.2 in the
**E2E Tests** workflow.

`run.sh` creates the kind cluster, installs Crossplane and the provider,
creates the ProviderConfigs, runs the Go tests in this folder, and then deletes
everything it created, in Infisical too. Every run has its own run ID, which is
part of the names of the Infisical resources, so runs in parallel do not
collide. If a resource cannot be deleted, the run prints a warning with the run
ID: delete the Infisical resources with that ID in their name by hand.

## Settings

| Environment variable | GitHub secret | Required | What it is |
|---|---|---|---|
| `INFISICAL_HOST` | `E2E_INFISICAL_HOST` | yes | URL of the Infisical instance, for example `https://app.infisical.com`. |
| `INFISICAL_CLIENT_ID` | `INFISICAL_CLIENT_ID` | yes | Client ID of a machine identity with Universal Auth. The provider authenticates with it. |
| `INFISICAL_CLIENT_SECRET` | `INFISICAL_CLIENT_SECRET` | yes | Client secret of that machine identity. |
| `INFISICAL_ORG_ID` | `INFISICAL_ORG_ID` | yes | ID of the organization in which the tests create machine identities. |
| `INFISICAL_USER_EMAIL` | `INFISICAL_USER_EMAIL` | yes | Email of an existing user of the organization. The tests add the user to a project and make them the approver of the approval policies. |
| `INFISICAL_GITHUB_CONNECTION_ID` | `INFISICAL_GITHUB_CONNECTION_ID` | no | ID of an existing GitHub App Connection in the organization. |
| `INFISICAL_GITHUB_REPO_OWNER` | `INFISICAL_GITHUB_REPO_OWNER` | no | Owner of the GitHub repository that the GitHub secret sync writes to. |
| `INFISICAL_GITHUB_REPO_NAME` | `INFISICAL_GITHUB_REPO_NAME` | no | Name of that repository. Use a repository for tests only: the sync writes secrets to it. |
| `INFISICAL_E2E_PAID_PLAN` | `INFISICAL_E2E_PAID_PLAN` | no | Set to `true` when the organization has a paid Infisical plan, to also test the kinds that need one. |

The tests create and delete projects, environments, folders, secrets, machine
identities and their auth methods, groups, project roles and memberships,
approval policies, project templates and secret syncs (see `objects_test.go`),
so the machine identity needs permission to manage these in the organization.

The three GitHub settings are optional. Unless all three are set, the
`SecretSyncGithub` test does not run, and the run prints a warning. The other
tests run as usual.

Some kinds need a paid Infisical plan. On the free plan, Infisical rejects
groups, secret approval policies and project templates, and allows three
environments per project, which a new project already has. Unless
`INFISICAL_E2E_PAID_PLAN` is `true`, the `Group`, `ProjectGroup`,
`ProjectEnvironment`, `SecretApprovalPolicy` and `ProjectTemplate` tests do not
run, and the run prints a warning. To see what the plan of an organization
allows, call `GET /api/v1/organizations/<org ID>/plan`.

The workflow takes the settings from the repository secrets in the last
column. Note that the host comes from the secret `E2E_INFISICAL_HOST`.

## Run locally

You need Docker, Go and `jq`. The Makefile installs kind, Helm, kubectl and
the Crossplane CLI.

1. Put the settings into a file, for example `e2e.env` (do not commit it):

   ```sh
   INFISICAL_HOST=https://app.infisical.com
   INFISICAL_CLIENT_ID=...
   INFISICAL_CLIENT_SECRET=...
   INFISICAL_ORG_ID=...
   INFISICAL_USER_EMAIL=...
   # Optional:
   # INFISICAL_GITHUB_CONNECTION_ID=...
   # INFISICAL_GITHUB_REPO_OWNER=...
   # INFISICAL_GITHUB_REPO_NAME=...
   ```

2. Build the provider image and package, then run the tests:

   ```sh
   make build
   E2E_ENV_FILE=$PWD/e2e.env make test-e2e CROSSPLANE_VERSION=2.4.2
   ```

`KEEP_CLUSTER=true` keeps the kind cluster after the run. The test results
(JSON and JUnit) and, after a failure, the diagnostics are written to
`_output/e2e/`.

## Add a test

Add a cluster scoped fixture to `testdata/` and an entry to `objects_test.go`.
The tests render the namespaced variant from the same fixture.
