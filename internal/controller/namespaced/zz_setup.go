// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appconnectionazureclientsecrets "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionazureclientsecrets"
	group "github.com/infisical/provider-infisical/internal/controller/namespaced/group/group"
	identity "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/identity"
	kubernetesauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/kubernetesauth"
	universalauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/universalauth"
	accessapprovalpolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/project/accessapprovalpolicy"
	project "github.com/infisical/provider-infisical/internal/controller/namespaced/project/project"
	projectenvironment "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectenvironment"
	projectgroup "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectgroup"
	projectidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectidentity"
	projectrole "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectrole"
	projecttemplate "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projecttemplate"
	projectuser "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectuser"
	secretapprovalpolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/project/secretapprovalpolicy"
	providerconfig "github.com/infisical/provider-infisical/internal/controller/namespaced/providerconfig"
	secret "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secret"
	secretfolder "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretfolder"
	secretrotationazureclientsecret "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationazureclientsecret"
	secretsyncgithub "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncgithub"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnectionazureclientsecrets.Setup,
		group.Setup,
		identity.Setup,
		kubernetesauth.Setup,
		universalauth.Setup,
		accessapprovalpolicy.Setup,
		project.Setup,
		projectenvironment.Setup,
		projectgroup.Setup,
		projectidentity.Setup,
		projectrole.Setup,
		projecttemplate.Setup,
		projectuser.Setup,
		secretapprovalpolicy.Setup,
		providerconfig.Setup,
		secret.Setup,
		secretfolder.Setup,
		secretrotationazureclientsecret.Setup,
		secretsyncgithub.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnectionazureclientsecrets.SetupGated,
		group.SetupGated,
		identity.SetupGated,
		kubernetesauth.SetupGated,
		universalauth.SetupGated,
		accessapprovalpolicy.SetupGated,
		project.SetupGated,
		projectenvironment.SetupGated,
		projectgroup.SetupGated,
		projectidentity.SetupGated,
		projectrole.SetupGated,
		projecttemplate.SetupGated,
		projectuser.SetupGated,
		secretapprovalpolicy.SetupGated,
		providerconfig.SetupGated,
		secret.SetupGated,
		secretfolder.SetupGated,
		secretrotationazureclientsecret.SetupGated,
		secretsyncgithub.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		appconnectionazureclientsecrets.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		identity.SetupWebhookWithManager,
		kubernetesauth.SetupWebhookWithManager,
		universalauth.SetupWebhookWithManager,
		accessapprovalpolicy.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		projectenvironment.SetupWebhookWithManager,
		projectgroup.SetupWebhookWithManager,
		projectidentity.SetupWebhookWithManager,
		projectrole.SetupWebhookWithManager,
		projecttemplate.SetupWebhookWithManager,
		projectuser.SetupWebhookWithManager,
		secretapprovalpolicy.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
		secretfolder.SetupWebhookWithManager,
		secretrotationazureclientsecret.SetupWebhookWithManager,
		secretsyncgithub.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
