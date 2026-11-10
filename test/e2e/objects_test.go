//go:build e2e

/*
Copyright 2026 Infisical Inc.
*/

package e2e

import (
	"fmt"
	"os"
	"path/filepath"
)

// object is one managed resource of the test fixtures in testdata. The
// fixtures are cluster scoped, and the test derives the namespaced object
// from the same fixture.
type object struct {
	// file is the name of the fixture file, without ".yaml".
	file string
	// name is the suffix of the Kubernetes name, after the set prefix.
	name string
	// shortGroup is the API group without the root group, for example
	// "project".
	shortGroup string
	kind       string

	// parent objects are created first and deleted last, because other
	// objects reference them.
	parent bool

	// notReady is why the object cannot become ready in the test. "" means
	// it must become ready.
	notReady string

	// paidPlan is the feature of a paid Infisical plan that the object needs.
	// "" means it works on the free plan. Objects that need a paid plan are
	// only tested when INFISICAL_E2E_PAID_PLAN is "true".
	paidPlan string

	// update is the update test: a merge patch of spec.forProvider, and the
	// value that status.atProvider must then have. Values are JSON with the
	// placeholders of the fixtures.
	update *update
}

type update struct {
	set   string
	check string
	want  string
}

// testName returns the name of the object in test output.
func (o object) testName() string {
	if o.file == "identity-kubernetes" {
		return "Identity(for-KubernetesAuth)"
	}
	return o.kind
}

var (
	reasonKubernetesAuth = "Infisical calls the Kubernetes API of kubernetesHost when the auth method is created, and the test has no Kubernetes API that Infisical can reach"
	reasonProjectRole    = "known gap: the ProjectRole Read of the Terraform provider returns an error instead of removing the resource when the role is not found, so a new ProjectRole is never created"
)

// The features of a paid Infisical plan that test objects need. The free plan
// has no groups, secret approval policies or project templates, and allows
// three environments per project, which a new project already has.
const (
	paidGroups           = "groups"
	paidEnvironmentLimit = "more than three environments per project"
	paidSecretApproval   = "secret approval policies"
	paidProjectTemplates = "project templates"
)

func updateOf(set, check, want string) *update {
	return &update{set: set, check: check, want: want}
}

const (
	projectGroup    = "project"
	identityGroup   = "identity"
	secretGroup     = "secret"
	groupGroup      = "group"
	secretSyncGroup = "secretsync"
)

// objects are all test objects. Every kind of the provider has at least one.
var objects = []object{
	{
		file: "project", name: "project", shortGroup: projectGroup, kind: "Project", parent: true,
		update: updateOf(`{"description":"updated by e2e"}`, "status.atProvider.description", `"updated by e2e"`),
	},
	{
		file: "identity", name: "identity", shortGroup: identityGroup, kind: "Identity", parent: true,
		update: updateOf(`{"role":"no-access"}`, "status.atProvider.role", `"no-access"`),
	},
	{
		file: "identity-kubernetes", name: "identity-kubernetes", shortGroup: identityGroup, kind: "Identity", parent: true,
	},
	{
		file: "group", name: "group", shortGroup: groupGroup, kind: "Group", parent: true, paidPlan: paidGroups,
		update: updateOf(`{"name":"__PREFIX__-__RUN_ID__-updated"}`, "status.atProvider.name", `"__PREFIX__-__RUN_ID__-updated"`),
	},
	{
		file: "projectenvironment", name: "environment", shortGroup: projectGroup, kind: "ProjectEnvironment", paidPlan: paidEnvironmentLimit,
		update: updateOf(`{"name":"e2e-updated"}`, "status.atProvider.name", `"e2e-updated"`),
	},
	{
		file: "secretfolder", name: "folder", shortGroup: secretGroup, kind: "SecretFolder",
		update: updateOf(`{"description":"updated by e2e"}`, "status.atProvider.description", `"updated by e2e"`),
	},
	{
		file: "universalauth", name: "universal-auth", shortGroup: identityGroup, kind: "UniversalAuth",
		update: updateOf(`{"accessTokenTtl":3600}`, "status.atProvider.accessTokenTtl", `3600`),
	},
	{
		file: "kubernetesauth", name: "kubernetes-auth", shortGroup: identityGroup, kind: "KubernetesAuth",
		notReady: reasonKubernetesAuth,
	},
	{
		file: "projectidentity", name: "project-identity", shortGroup: projectGroup, kind: "ProjectIdentity",
		update: updateOf(`{"roles":[{"roleSlug":"member"}]}`, "status.atProvider.roles", `[{"roleSlug":"member"}]`),
	},
	{
		file: "projectuser", name: "project-user", shortGroup: projectGroup, kind: "ProjectUser",
		update: updateOf(`{"roles":[{"roleSlug":"member"}]}`, "status.atProvider.roles", `[{"roleSlug":"member"}]`),
	},
	{
		file: "projectgroup", name: "project-group", shortGroup: projectGroup, kind: "ProjectGroup", paidPlan: paidGroups,
		update: updateOf(`{"roles":[{"roleSlug":"member"}]}`, "status.atProvider.roles", `[{"roleSlug":"member"}]`),
	},
	{
		file: "projectrole", name: "project-role", shortGroup: projectGroup, kind: "ProjectRole",
		notReady: reasonProjectRole,
		update:   updateOf(`{"description":"updated by e2e"}`, "status.atProvider.description", `"updated by e2e"`),
	},
	{
		file: "accessapprovalpolicy", name: "access-approval-policy", shortGroup: projectGroup, kind: "AccessApprovalPolicy",
		update: updateOf(`{"enforcementLevel":"hard"}`, "status.atProvider.enforcementLevel", `"hard"`),
	},
	{
		file: "secretapprovalpolicy", name: "secret-approval-policy", shortGroup: projectGroup, kind: "SecretApprovalPolicy", paidPlan: paidSecretApproval,
		update: updateOf(`{"enforcementLevel":"hard"}`, "status.atProvider.enforcementLevel", `"hard"`),
	},
	{
		file: "projecttemplate", name: "project-template", shortGroup: projectGroup, kind: "ProjectTemplate", paidPlan: paidProjectTemplates,
		update: updateOf(`{"description":"updated by e2e"}`, "status.atProvider.description", `"updated by e2e"`),
	},
	{
		file: githubSecretSyncFile, name: "secret-sync-github", shortGroup: secretSyncGroup, kind: "SecretSyncGithub",
		update: updateOf(`{"syncOptions":{"keySchema":"E2E_UPDATED_{{secretKey}}"}}`, "status.atProvider.syncOptions.keySchema", `"E2E_UPDATED_{{secretKey}}"`),
	},
	{
		file: "secret", name: "secret", shortGroup: secretGroup, kind: "Secret",
		update: updateOf(`{"secretReminder":{"note":"updated by e2e"}}`, "status.atProvider.secretReminder.note", `"updated by e2e"`),
	},
}

// checkFixtures returns the problems with the fixture files: every object
// needs a fixture.
func checkFixtures() []string {
	var problems []string
	for _, o := range objects {
		path := filepath.Join("testdata", o.file+".yaml")
		if _, err := os.Stat(path); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", path, err))
		}
	}
	return problems
}

// githubSecretSyncFile is the fixture of the GitHub secret sync. Its test only
// runs when the optional GitHub settings are set.
const githubSecretSyncFile = "secretsyncgithub"

// withoutObject returns the test objects without the object of the fixture
// file.
func withoutObject(all []object, file string) []object {
	kept := make([]object, 0, len(all))
	for _, o := range all {
		if o.file != file {
			kept = append(kept, o)
		}
	}
	return kept
}

// withoutPaidPlan returns the test objects that work on the free Infisical
// plan, and the test names of the objects that it leaves out.
func withoutPaidPlan(all []object) ([]object, []string) {
	kept := make([]object, 0, len(all))
	var skippedNames []string
	for _, o := range all {
		if o.paidPlan == "" {
			kept = append(kept, o)
			continue
		}
		skippedNames = append(skippedNames, fmt.Sprintf("%s (%s)", o.testName(), o.paidPlan))
	}
	return kept, skippedNames
}

// objectByFile returns the test object of the fixture file.
func objectByFile(file string) object {
	for _, o := range objects {
		if o.file == file {
			return o
		}
	}
	panic("unknown test object " + file)
}
