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
		file: "group", name: "group", shortGroup: groupGroup, kind: "Group", parent: true,
		update: updateOf(`{"name":"__PREFIX__-__RUN_ID__-updated"}`, "status.atProvider.name", `"__PREFIX__-__RUN_ID__-updated"`),
	},
	{
		file: "projectenvironment", name: "environment", shortGroup: projectGroup, kind: "ProjectEnvironment",
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
		file: "projectgroup", name: "project-group", shortGroup: projectGroup, kind: "ProjectGroup",
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
		file: "secretapprovalpolicy", name: "secret-approval-policy", shortGroup: projectGroup, kind: "SecretApprovalPolicy",
		update: updateOf(`{"enforcementLevel":"hard"}`, "status.atProvider.enforcementLevel", `"hard"`),
	},
	{
		file: "projecttemplate", name: "project-template", shortGroup: projectGroup, kind: "ProjectTemplate",
		update: updateOf(`{"description":"updated by e2e"}`, "status.atProvider.description", `"updated by e2e"`),
	},
	{
		file: "secretsyncgithub", name: "secret-sync-github", shortGroup: secretSyncGroup, kind: "SecretSyncGithub",
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

// objectByFile returns the test object of the fixture file.
func objectByFile(file string) object {
	for _, o := range objects {
		if o.file == file {
			return o
		}
	}
	panic("unknown test object " + file)
}
