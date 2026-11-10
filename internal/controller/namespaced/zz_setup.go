// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appconnectionaws "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionaws"
	appconnectionazureappconfiguration "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionazureappconfiguration"
	appconnectionazureclientsecrets "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionazureclientsecrets"
	appconnectionazuredevops "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionazuredevops"
	appconnectionazurekeyvault "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionazurekeyvault"
	appconnectionbitbucket "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionbitbucket"
	appconnectionchef "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionchef"
	appconnectioncircleci "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectioncircleci"
	appconnectioncloudflare "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectioncloudflare"
	appconnectiondatabricks "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiondatabricks"
	appconnectiondatadog "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiondatadog"
	appconnectiondigicert "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiondigicert"
	appconnectionf5bigip "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionf5bigip"
	appconnectionflyio "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionflyio"
	appconnectiongcp "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiongcp"
	appconnectiongithub "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiongithub"
	appconnectiongitlab "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiongitlab"
	appconnectionhashicorpvault "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionhashicorpvault"
	appconnectionkemploadmaster "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionkemploadmaster"
	appconnectionldap "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionldap"
	appconnectionlitellm "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionlitellm"
	appconnectionmssql "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionmssql"
	appconnectionmysql "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionmysql"
	appconnectionnetscaler "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionnetscaler"
	appconnectionnutanixprismcentral "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionnutanixprismcentral"
	appconnectiononepassword "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectiononepassword"
	appconnectionoracledb "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionoracledb"
	appconnectionpostgres "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionpostgres"
	appconnectionrender "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionrender"
	appconnectionspacelift "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionspacelift"
	appconnectionssh "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionssh"
	appconnectionsupabase "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionsupabase"
	appconnectionwinrm "github.com/infisical/provider-infisical/internal/controller/namespaced/appconnection/appconnectionwinrm"
	awscertificatemanager "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/awscertificatemanager"
	awselasticloadbalancer "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/awselasticloadbalancer"
	awssecretsmanager "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/awssecretsmanager"
	azurekeyvault "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/azurekeyvault"
	certificate "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/certificate"
	chef "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/chef"
	cloudflarecustomcertificate "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/cloudflarecustomcertificate"
	f5bigip "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/f5bigip"
	gcpcertificatemanager "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/gcpcertificatemanager"
	kemploadmaster "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/kemploadmaster"
	linuxserver "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/linuxserver"
	netscaler "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/netscaler"
	nutanixprismcentral "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/nutanixprismcentral"
	windowsserver "github.com/infisical/provider-infisical/internal/controller/namespaced/certificatesync/windowsserver"
	application "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/application"
	applicationapprovalpolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/applicationapprovalpolicy"
	applicationgroup "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/applicationgroup"
	applicationidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/applicationidentity"
	applicationprofile "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/applicationprofile"
	applicationuser "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/applicationuser"
	cacertificate "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/cacertificate"
	certificatecertmanager "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/certificate"
	certificatepolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/certificatepolicy"
	certificateprofile "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/certificateprofile"
	externalcaacme "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/externalcaacme"
	externalcaadcs "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/externalcaadcs"
	externalcadigicert "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/externalcadigicert"
	group "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/group"
	identity "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/identity"
	internalca "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/internalca"
	user "github.com/infisical/provider-infisical/internal/controller/namespaced/certmanager/user"
	awsiam "github.com/infisical/provider-infisical/internal/controller/namespaced/dynamicsecret/awsiam"
	kubernetes "github.com/infisical/provider-infisical/internal/controller/namespaced/dynamicsecret/kubernetes"
	mongoatlas "github.com/infisical/provider-infisical/internal/controller/namespaced/dynamicsecret/mongoatlas"
	mongodb "github.com/infisical/provider-infisical/internal/controller/namespaced/dynamicsecret/mongodb"
	sqldatabase "github.com/infisical/provider-infisical/internal/controller/namespaced/dynamicsecret/sqldatabase"
	enrollmenttoken "github.com/infisical/provider-infisical/internal/controller/namespaced/gateway/enrollmenttoken"
	gateway "github.com/infisical/provider-infisical/internal/controller/namespaced/gateway/gateway"
	groupgroup "github.com/infisical/provider-infisical/internal/controller/namespaced/group/group"
	groupmachineidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/group/groupmachineidentity"
	awsauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/awsauth"
	azureauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/azureauth"
	gcpauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/gcpauth"
	identityidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/identity"
	jwtauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/jwtauth"
	kubernetesauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/kubernetesauth"
	kubernetesauthtemplate "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/kubernetesauthtemplate"
	oidcauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/oidcauth"
	oidcauthtemplate "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/oidcauthtemplate"
	tlscertauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/tlscertauth"
	tokenauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/tokenauth"
	tokenauthtoken "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/tokenauthtoken"
	universalauth "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/universalauth"
	universalauthclientsecret "github.com/infisical/provider-infisical/internal/controller/namespaced/identity/universalauthclientsecret"
	awsparameterstore "github.com/infisical/provider-infisical/internal/controller/namespaced/integration/awsparameterstore"
	awssecretsmanagerintegration "github.com/infisical/provider-infisical/internal/controller/namespaced/integration/awssecretsmanager"
	circleci "github.com/infisical/provider-infisical/internal/controller/namespaced/integration/circleci"
	databricks "github.com/infisical/provider-infisical/internal/controller/namespaced/integration/databricks"
	gcpsecretmanager "github.com/infisical/provider-infisical/internal/controller/namespaced/integration/gcpsecretmanager"
	externalaws "github.com/infisical/provider-infisical/internal/controller/namespaced/kms/externalaws"
	key "github.com/infisical/provider-infisical/internal/controller/namespaced/kms/key"
	alert "github.com/infisical/provider-infisical/internal/controller/namespaced/organization/alert"
	role "github.com/infisical/provider-infisical/internal/controller/namespaced/organization/role"
	suborganization "github.com/infisical/provider-infisical/internal/controller/namespaced/organization/suborganization"
	suborganizationgroup "github.com/infisical/provider-infisical/internal/controller/namespaced/organization/suborganizationgroup"
	accessapprovalpolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/project/accessapprovalpolicy"
	project "github.com/infisical/provider-infisical/internal/controller/namespaced/project/project"
	projectenvironment "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectenvironment"
	projectgroup "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectgroup"
	projectidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectidentity"
	projectidentityspecificprivilege "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectidentityspecificprivilege"
	projectrole "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectrole"
	projectscopedidentity "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectscopedidentity"
	projecttemplate "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projecttemplate"
	projectuser "github.com/infisical/provider-infisical/internal/controller/namespaced/project/projectuser"
	secretapprovalpolicy "github.com/infisical/provider-infisical/internal/controller/namespaced/project/secretapprovalpolicy"
	webhook "github.com/infisical/provider-infisical/internal/controller/namespaced/project/webhook"
	providerconfig "github.com/infisical/provider-infisical/internal/controller/namespaced/providerconfig"
	secret "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secret"
	secretfolder "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretfolder"
	secretimport "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretimport"
	secrettag "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secrettag"
	secretvalidationruledynamicsecrets "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretvalidationruledynamicsecrets"
	secretvalidationrulesecretrotations "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretvalidationrulesecretrotations"
	secretvalidationrulestaticsecrets "github.com/infisical/provider-infisical/internal/controller/namespaced/secret/secretvalidationrulestaticsecrets"
	secretrotationauth0clientsecret "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationauth0clientsecret"
	secretrotationawsiamusersecret "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationawsiamusersecret"
	secretrotationazureclientsecret "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationazureclientsecret"
	secretrotationdatadogapikey "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationdatadogapikey"
	secretrotationldappassword "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationldappassword"
	secretrotationlitellmapikey "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationlitellmapikey"
	secretrotationmssqlcredentials "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationmssqlcredentials"
	secretrotationmysqlcredentials "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationmysqlcredentials"
	secretrotationoracledbcredentials "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationoracledbcredentials"
	secretrotationpostgrescredentials "github.com/infisical/provider-infisical/internal/controller/namespaced/secretrotation/secretrotationpostgrescredentials"
	secretsyncawsparameterstore "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncawsparameterstore"
	secretsyncawssecretsmanager "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncawssecretsmanager"
	secretsyncazureappconfiguration "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncazureappconfiguration"
	secretsyncazuredevops "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncazuredevops"
	secretsyncazurekeyvault "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncazurekeyvault"
	secretsyncbitbucket "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncbitbucket"
	secretsynccircleci "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsynccircleci"
	secretsynccloudflarepages "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsynccloudflarepages"
	secretsynccloudflareworkers "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsynccloudflareworkers"
	secretsyncdatabricks "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncdatabricks"
	secretsyncflyio "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncflyio"
	secretsyncgcpsecretmanager "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncgcpsecretmanager"
	secretsyncgithub "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncgithub"
	secretsyncgitlab "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncgitlab"
	secretsynconepassword "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsynconepassword"
	secretsyncrender "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncrender"
	secretsyncspacelift "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncspacelift"
	secretsyncsupabase "github.com/infisical/provider-infisical/internal/controller/namespaced/secretsync/secretsyncsupabase"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnectionaws.Setup,
		appconnectionazureappconfiguration.Setup,
		appconnectionazureclientsecrets.Setup,
		appconnectionazuredevops.Setup,
		appconnectionazurekeyvault.Setup,
		appconnectionbitbucket.Setup,
		appconnectionchef.Setup,
		appconnectioncircleci.Setup,
		appconnectioncloudflare.Setup,
		appconnectiondatabricks.Setup,
		appconnectiondatadog.Setup,
		appconnectiondigicert.Setup,
		appconnectionf5bigip.Setup,
		appconnectionflyio.Setup,
		appconnectiongcp.Setup,
		appconnectiongithub.Setup,
		appconnectiongitlab.Setup,
		appconnectionhashicorpvault.Setup,
		appconnectionkemploadmaster.Setup,
		appconnectionldap.Setup,
		appconnectionlitellm.Setup,
		appconnectionmssql.Setup,
		appconnectionmysql.Setup,
		appconnectionnetscaler.Setup,
		appconnectionnutanixprismcentral.Setup,
		appconnectiononepassword.Setup,
		appconnectionoracledb.Setup,
		appconnectionpostgres.Setup,
		appconnectionrender.Setup,
		appconnectionspacelift.Setup,
		appconnectionssh.Setup,
		appconnectionsupabase.Setup,
		appconnectionwinrm.Setup,
		awscertificatemanager.Setup,
		awselasticloadbalancer.Setup,
		awssecretsmanager.Setup,
		azurekeyvault.Setup,
		certificate.Setup,
		chef.Setup,
		cloudflarecustomcertificate.Setup,
		f5bigip.Setup,
		gcpcertificatemanager.Setup,
		kemploadmaster.Setup,
		linuxserver.Setup,
		netscaler.Setup,
		nutanixprismcentral.Setup,
		windowsserver.Setup,
		application.Setup,
		applicationapprovalpolicy.Setup,
		applicationgroup.Setup,
		applicationidentity.Setup,
		applicationprofile.Setup,
		applicationuser.Setup,
		cacertificate.Setup,
		certificatecertmanager.Setup,
		certificatepolicy.Setup,
		certificateprofile.Setup,
		externalcaacme.Setup,
		externalcaadcs.Setup,
		externalcadigicert.Setup,
		group.Setup,
		identity.Setup,
		internalca.Setup,
		user.Setup,
		awsiam.Setup,
		kubernetes.Setup,
		mongoatlas.Setup,
		mongodb.Setup,
		sqldatabase.Setup,
		enrollmenttoken.Setup,
		gateway.Setup,
		groupgroup.Setup,
		groupmachineidentity.Setup,
		awsauth.Setup,
		azureauth.Setup,
		gcpauth.Setup,
		identityidentity.Setup,
		jwtauth.Setup,
		kubernetesauth.Setup,
		kubernetesauthtemplate.Setup,
		oidcauth.Setup,
		oidcauthtemplate.Setup,
		tlscertauth.Setup,
		tokenauth.Setup,
		tokenauthtoken.Setup,
		universalauth.Setup,
		universalauthclientsecret.Setup,
		awsparameterstore.Setup,
		awssecretsmanagerintegration.Setup,
		circleci.Setup,
		databricks.Setup,
		gcpsecretmanager.Setup,
		externalaws.Setup,
		key.Setup,
		alert.Setup,
		role.Setup,
		suborganization.Setup,
		suborganizationgroup.Setup,
		accessapprovalpolicy.Setup,
		project.Setup,
		projectenvironment.Setup,
		projectgroup.Setup,
		projectidentity.Setup,
		projectidentityspecificprivilege.Setup,
		projectrole.Setup,
		projectscopedidentity.Setup,
		projecttemplate.Setup,
		projectuser.Setup,
		secretapprovalpolicy.Setup,
		webhook.Setup,
		providerconfig.Setup,
		secret.Setup,
		secretfolder.Setup,
		secretimport.Setup,
		secrettag.Setup,
		secretvalidationruledynamicsecrets.Setup,
		secretvalidationrulesecretrotations.Setup,
		secretvalidationrulestaticsecrets.Setup,
		secretrotationauth0clientsecret.Setup,
		secretrotationawsiamusersecret.Setup,
		secretrotationazureclientsecret.Setup,
		secretrotationdatadogapikey.Setup,
		secretrotationldappassword.Setup,
		secretrotationlitellmapikey.Setup,
		secretrotationmssqlcredentials.Setup,
		secretrotationmysqlcredentials.Setup,
		secretrotationoracledbcredentials.Setup,
		secretrotationpostgrescredentials.Setup,
		secretsyncawsparameterstore.Setup,
		secretsyncawssecretsmanager.Setup,
		secretsyncazureappconfiguration.Setup,
		secretsyncazuredevops.Setup,
		secretsyncazurekeyvault.Setup,
		secretsyncbitbucket.Setup,
		secretsynccircleci.Setup,
		secretsynccloudflarepages.Setup,
		secretsynccloudflareworkers.Setup,
		secretsyncdatabricks.Setup,
		secretsyncflyio.Setup,
		secretsyncgcpsecretmanager.Setup,
		secretsyncgithub.Setup,
		secretsyncgitlab.Setup,
		secretsynconepassword.Setup,
		secretsyncrender.Setup,
		secretsyncspacelift.Setup,
		secretsyncsupabase.Setup,
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
		appconnectionaws.SetupGated,
		appconnectionazureappconfiguration.SetupGated,
		appconnectionazureclientsecrets.SetupGated,
		appconnectionazuredevops.SetupGated,
		appconnectionazurekeyvault.SetupGated,
		appconnectionbitbucket.SetupGated,
		appconnectionchef.SetupGated,
		appconnectioncircleci.SetupGated,
		appconnectioncloudflare.SetupGated,
		appconnectiondatabricks.SetupGated,
		appconnectiondatadog.SetupGated,
		appconnectiondigicert.SetupGated,
		appconnectionf5bigip.SetupGated,
		appconnectionflyio.SetupGated,
		appconnectiongcp.SetupGated,
		appconnectiongithub.SetupGated,
		appconnectiongitlab.SetupGated,
		appconnectionhashicorpvault.SetupGated,
		appconnectionkemploadmaster.SetupGated,
		appconnectionldap.SetupGated,
		appconnectionlitellm.SetupGated,
		appconnectionmssql.SetupGated,
		appconnectionmysql.SetupGated,
		appconnectionnetscaler.SetupGated,
		appconnectionnutanixprismcentral.SetupGated,
		appconnectiononepassword.SetupGated,
		appconnectionoracledb.SetupGated,
		appconnectionpostgres.SetupGated,
		appconnectionrender.SetupGated,
		appconnectionspacelift.SetupGated,
		appconnectionssh.SetupGated,
		appconnectionsupabase.SetupGated,
		appconnectionwinrm.SetupGated,
		awscertificatemanager.SetupGated,
		awselasticloadbalancer.SetupGated,
		awssecretsmanager.SetupGated,
		azurekeyvault.SetupGated,
		certificate.SetupGated,
		chef.SetupGated,
		cloudflarecustomcertificate.SetupGated,
		f5bigip.SetupGated,
		gcpcertificatemanager.SetupGated,
		kemploadmaster.SetupGated,
		linuxserver.SetupGated,
		netscaler.SetupGated,
		nutanixprismcentral.SetupGated,
		windowsserver.SetupGated,
		application.SetupGated,
		applicationapprovalpolicy.SetupGated,
		applicationgroup.SetupGated,
		applicationidentity.SetupGated,
		applicationprofile.SetupGated,
		applicationuser.SetupGated,
		cacertificate.SetupGated,
		certificatecertmanager.SetupGated,
		certificatepolicy.SetupGated,
		certificateprofile.SetupGated,
		externalcaacme.SetupGated,
		externalcaadcs.SetupGated,
		externalcadigicert.SetupGated,
		group.SetupGated,
		identity.SetupGated,
		internalca.SetupGated,
		user.SetupGated,
		awsiam.SetupGated,
		kubernetes.SetupGated,
		mongoatlas.SetupGated,
		mongodb.SetupGated,
		sqldatabase.SetupGated,
		enrollmenttoken.SetupGated,
		gateway.SetupGated,
		groupgroup.SetupGated,
		groupmachineidentity.SetupGated,
		awsauth.SetupGated,
		azureauth.SetupGated,
		gcpauth.SetupGated,
		identityidentity.SetupGated,
		jwtauth.SetupGated,
		kubernetesauth.SetupGated,
		kubernetesauthtemplate.SetupGated,
		oidcauth.SetupGated,
		oidcauthtemplate.SetupGated,
		tlscertauth.SetupGated,
		tokenauth.SetupGated,
		tokenauthtoken.SetupGated,
		universalauth.SetupGated,
		universalauthclientsecret.SetupGated,
		awsparameterstore.SetupGated,
		awssecretsmanagerintegration.SetupGated,
		circleci.SetupGated,
		databricks.SetupGated,
		gcpsecretmanager.SetupGated,
		externalaws.SetupGated,
		key.SetupGated,
		alert.SetupGated,
		role.SetupGated,
		suborganization.SetupGated,
		suborganizationgroup.SetupGated,
		accessapprovalpolicy.SetupGated,
		project.SetupGated,
		projectenvironment.SetupGated,
		projectgroup.SetupGated,
		projectidentity.SetupGated,
		projectidentityspecificprivilege.SetupGated,
		projectrole.SetupGated,
		projectscopedidentity.SetupGated,
		projecttemplate.SetupGated,
		projectuser.SetupGated,
		secretapprovalpolicy.SetupGated,
		webhook.SetupGated,
		providerconfig.SetupGated,
		secret.SetupGated,
		secretfolder.SetupGated,
		secretimport.SetupGated,
		secrettag.SetupGated,
		secretvalidationruledynamicsecrets.SetupGated,
		secretvalidationrulesecretrotations.SetupGated,
		secretvalidationrulestaticsecrets.SetupGated,
		secretrotationauth0clientsecret.SetupGated,
		secretrotationawsiamusersecret.SetupGated,
		secretrotationazureclientsecret.SetupGated,
		secretrotationdatadogapikey.SetupGated,
		secretrotationldappassword.SetupGated,
		secretrotationlitellmapikey.SetupGated,
		secretrotationmssqlcredentials.SetupGated,
		secretrotationmysqlcredentials.SetupGated,
		secretrotationoracledbcredentials.SetupGated,
		secretrotationpostgrescredentials.SetupGated,
		secretsyncawsparameterstore.SetupGated,
		secretsyncawssecretsmanager.SetupGated,
		secretsyncazureappconfiguration.SetupGated,
		secretsyncazuredevops.SetupGated,
		secretsyncazurekeyvault.SetupGated,
		secretsyncbitbucket.SetupGated,
		secretsynccircleci.SetupGated,
		secretsynccloudflarepages.SetupGated,
		secretsynccloudflareworkers.SetupGated,
		secretsyncdatabricks.SetupGated,
		secretsyncflyio.SetupGated,
		secretsyncgcpsecretmanager.SetupGated,
		secretsyncgithub.SetupGated,
		secretsyncgitlab.SetupGated,
		secretsynconepassword.SetupGated,
		secretsyncrender.SetupGated,
		secretsyncspacelift.SetupGated,
		secretsyncsupabase.SetupGated,
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
		appconnectionaws.SetupWebhookWithManager,
		appconnectionazureappconfiguration.SetupWebhookWithManager,
		appconnectionazureclientsecrets.SetupWebhookWithManager,
		appconnectionazuredevops.SetupWebhookWithManager,
		appconnectionazurekeyvault.SetupWebhookWithManager,
		appconnectionbitbucket.SetupWebhookWithManager,
		appconnectionchef.SetupWebhookWithManager,
		appconnectioncircleci.SetupWebhookWithManager,
		appconnectioncloudflare.SetupWebhookWithManager,
		appconnectiondatabricks.SetupWebhookWithManager,
		appconnectiondatadog.SetupWebhookWithManager,
		appconnectiondigicert.SetupWebhookWithManager,
		appconnectionf5bigip.SetupWebhookWithManager,
		appconnectionflyio.SetupWebhookWithManager,
		appconnectiongcp.SetupWebhookWithManager,
		appconnectiongithub.SetupWebhookWithManager,
		appconnectiongitlab.SetupWebhookWithManager,
		appconnectionhashicorpvault.SetupWebhookWithManager,
		appconnectionkemploadmaster.SetupWebhookWithManager,
		appconnectionldap.SetupWebhookWithManager,
		appconnectionlitellm.SetupWebhookWithManager,
		appconnectionmssql.SetupWebhookWithManager,
		appconnectionmysql.SetupWebhookWithManager,
		appconnectionnetscaler.SetupWebhookWithManager,
		appconnectionnutanixprismcentral.SetupWebhookWithManager,
		appconnectiononepassword.SetupWebhookWithManager,
		appconnectionoracledb.SetupWebhookWithManager,
		appconnectionpostgres.SetupWebhookWithManager,
		appconnectionrender.SetupWebhookWithManager,
		appconnectionspacelift.SetupWebhookWithManager,
		appconnectionssh.SetupWebhookWithManager,
		appconnectionsupabase.SetupWebhookWithManager,
		appconnectionwinrm.SetupWebhookWithManager,
		awscertificatemanager.SetupWebhookWithManager,
		awselasticloadbalancer.SetupWebhookWithManager,
		awssecretsmanager.SetupWebhookWithManager,
		azurekeyvault.SetupWebhookWithManager,
		certificate.SetupWebhookWithManager,
		chef.SetupWebhookWithManager,
		cloudflarecustomcertificate.SetupWebhookWithManager,
		f5bigip.SetupWebhookWithManager,
		gcpcertificatemanager.SetupWebhookWithManager,
		kemploadmaster.SetupWebhookWithManager,
		linuxserver.SetupWebhookWithManager,
		netscaler.SetupWebhookWithManager,
		nutanixprismcentral.SetupWebhookWithManager,
		windowsserver.SetupWebhookWithManager,
		application.SetupWebhookWithManager,
		applicationapprovalpolicy.SetupWebhookWithManager,
		applicationgroup.SetupWebhookWithManager,
		applicationidentity.SetupWebhookWithManager,
		applicationprofile.SetupWebhookWithManager,
		applicationuser.SetupWebhookWithManager,
		cacertificate.SetupWebhookWithManager,
		certificatecertmanager.SetupWebhookWithManager,
		certificatepolicy.SetupWebhookWithManager,
		certificateprofile.SetupWebhookWithManager,
		externalcaacme.SetupWebhookWithManager,
		externalcaadcs.SetupWebhookWithManager,
		externalcadigicert.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		identity.SetupWebhookWithManager,
		internalca.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		awsiam.SetupWebhookWithManager,
		kubernetes.SetupWebhookWithManager,
		mongoatlas.SetupWebhookWithManager,
		mongodb.SetupWebhookWithManager,
		sqldatabase.SetupWebhookWithManager,
		enrollmenttoken.SetupWebhookWithManager,
		gateway.SetupWebhookWithManager,
		groupgroup.SetupWebhookWithManager,
		groupmachineidentity.SetupWebhookWithManager,
		awsauth.SetupWebhookWithManager,
		azureauth.SetupWebhookWithManager,
		gcpauth.SetupWebhookWithManager,
		identityidentity.SetupWebhookWithManager,
		jwtauth.SetupWebhookWithManager,
		kubernetesauth.SetupWebhookWithManager,
		kubernetesauthtemplate.SetupWebhookWithManager,
		oidcauth.SetupWebhookWithManager,
		oidcauthtemplate.SetupWebhookWithManager,
		tlscertauth.SetupWebhookWithManager,
		tokenauth.SetupWebhookWithManager,
		tokenauthtoken.SetupWebhookWithManager,
		universalauth.SetupWebhookWithManager,
		universalauthclientsecret.SetupWebhookWithManager,
		awsparameterstore.SetupWebhookWithManager,
		awssecretsmanagerintegration.SetupWebhookWithManager,
		circleci.SetupWebhookWithManager,
		databricks.SetupWebhookWithManager,
		gcpsecretmanager.SetupWebhookWithManager,
		externalaws.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
		alert.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		suborganization.SetupWebhookWithManager,
		suborganizationgroup.SetupWebhookWithManager,
		accessapprovalpolicy.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		projectenvironment.SetupWebhookWithManager,
		projectgroup.SetupWebhookWithManager,
		projectidentity.SetupWebhookWithManager,
		projectidentityspecificprivilege.SetupWebhookWithManager,
		projectrole.SetupWebhookWithManager,
		projectscopedidentity.SetupWebhookWithManager,
		projecttemplate.SetupWebhookWithManager,
		projectuser.SetupWebhookWithManager,
		secretapprovalpolicy.SetupWebhookWithManager,
		webhook.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
		secretfolder.SetupWebhookWithManager,
		secretimport.SetupWebhookWithManager,
		secrettag.SetupWebhookWithManager,
		secretvalidationruledynamicsecrets.SetupWebhookWithManager,
		secretvalidationrulesecretrotations.SetupWebhookWithManager,
		secretvalidationrulestaticsecrets.SetupWebhookWithManager,
		secretrotationauth0clientsecret.SetupWebhookWithManager,
		secretrotationawsiamusersecret.SetupWebhookWithManager,
		secretrotationazureclientsecret.SetupWebhookWithManager,
		secretrotationdatadogapikey.SetupWebhookWithManager,
		secretrotationldappassword.SetupWebhookWithManager,
		secretrotationlitellmapikey.SetupWebhookWithManager,
		secretrotationmssqlcredentials.SetupWebhookWithManager,
		secretrotationmysqlcredentials.SetupWebhookWithManager,
		secretrotationoracledbcredentials.SetupWebhookWithManager,
		secretrotationpostgrescredentials.SetupWebhookWithManager,
		secretsyncawsparameterstore.SetupWebhookWithManager,
		secretsyncawssecretsmanager.SetupWebhookWithManager,
		secretsyncazureappconfiguration.SetupWebhookWithManager,
		secretsyncazuredevops.SetupWebhookWithManager,
		secretsyncazurekeyvault.SetupWebhookWithManager,
		secretsyncbitbucket.SetupWebhookWithManager,
		secretsynccircleci.SetupWebhookWithManager,
		secretsynccloudflarepages.SetupWebhookWithManager,
		secretsynccloudflareworkers.SetupWebhookWithManager,
		secretsyncdatabricks.SetupWebhookWithManager,
		secretsyncflyio.SetupWebhookWithManager,
		secretsyncgcpsecretmanager.SetupWebhookWithManager,
		secretsyncgithub.SetupWebhookWithManager,
		secretsyncgitlab.SetupWebhookWithManager,
		secretsynconepassword.SetupWebhookWithManager,
		secretsyncrender.SetupWebhookWithManager,
		secretsyncspacelift.SetupWebhookWithManager,
		secretsyncsupabase.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
