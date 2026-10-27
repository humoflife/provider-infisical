# Provider Infisical

`provider-infisical` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Infisical API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/infisical-inc/provider-infisical):

```
up ctp provider install xpkg.upbound.io/infisical-inc/provider-infisical:v0.1.4
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-infisical
spec:
  package: xpkg.upbound.io/infisical-inc/provider-infisical:v0.1.4
EOF
```

### 🔐 Configuring Credentials

Before using the provider, create a Kubernetes `Secret` that contains your Infisical API credentials.

#### 1. Create a Secret

##### Authenticating with Universal Auth

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "auth": {
        "universal": {
          "client_id": "52eae513-7722-4bf0-af84-54cb1fa7d603",
          "client_secret": "672ae0e2a7b3ee97000d992ffc16726618cb3e4a43776978ca89f33e29ffdda0"
        }
      }
    }
```

##### Authenticating with Kubernetes Auth

This requires you to setup [Kubernetes Auth](https://infisical.com/docs/documentation/platform/identities/kubernetes-auth) with the controller's service account.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds-1
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "auth": {
        "kubernetes": {
          "identity_id": "81310d3c-aa51-4299-93df-3a7f983aee1a"
        }
      }
    }
```

##### Authenticating with Token Auth

This requires you to setup [Token Auth](https://infisical.com/docs/documentation/platform/identities/token-auth).

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds-1
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "auth": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV..."
      }
    }
```

#### 2. Create a ProviderConfig

```yaml
apiVersion: infisical.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  host: https://app.infisical.com
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: example-creds
      key: credentials
```

Now you can use `providerConfigRef.name: default` in your managed resources.

The cluster scoped managed resources are in the `*.infisical.upbound.io` API
groups, for example `project.infisical.upbound.io/v1beta1`.

#### Managed resources

The provider has a managed resource for each of the 145 resources of the
Infisical Terraform provider. Every kind exists cluster scoped and namespaced.

| API group | Kinds |
|---|---|
| `appconnection.infisical.upbound.io` | AppConnectionAws, AppConnectionAzureAppConfiguration, AppConnectionAzureClientSecrets, AppConnectionAzureDevops, AppConnectionAzureKeyVault, AppConnectionBitbucket, AppConnectionChef, AppConnectionCircleci, AppConnectionCloudflare, AppConnectionDatabricks, AppConnectionDatadog, AppConnectionDigicert, AppConnectionF5BigIp, AppConnectionFlyio, AppConnectionGcp, AppConnectionGithub, AppConnectionGitlab, AppConnectionHashicorpVault, AppConnectionKempLoadmaster, AppConnectionLdap, AppConnectionLitellm, AppConnectionMssql, AppConnectionMysql, AppConnectionNetscaler, AppConnectionNutanixPrismCentral, AppConnectionOnePassword, AppConnectionOracledb, AppConnectionPostgres, AppConnectionRender, AppConnectionSpacelift, AppConnectionSsh, AppConnectionSupabase, AppConnectionWinrm |
| `certificatesync.infisical.upbound.io` | AwsCertificateManager, AwsElasticLoadBalancer, AwsSecretsManager, AzureKeyVault, Certificate, Chef, CloudflareCustomCertificate, F5BigIp, GcpCertificateManager, KempLoadmaster, LinuxServer, Netscaler, NutanixPrismCentral, WindowsServer |
| `certmanager.infisical.upbound.io` | Application, ApplicationApprovalPolicy, ApplicationGroup, ApplicationIdentity, ApplicationProfile, ApplicationUser, CaCertificate, Certificate, CertificatePolicy, CertificateProfile, ExternalCaAcme, ExternalCaAdcs, ExternalCaDigicert, Group, Identity, InternalCa, User |
| `dynamicsecret.infisical.upbound.io` | AwsIam, Kubernetes, MongoAtlas, MongoDb, SqlDatabase |
| `gateway.infisical.upbound.io` | EnrollmentToken, Gateway |
| `group.infisical.upbound.io` | Group, GroupMachineIdentity |
| `identity.infisical.upbound.io` | AwsAuth, AzureAuth, GcpAuth, Identity, JwtAuth, KubernetesAuth, KubernetesAuthTemplate, OidcAuth, OidcAuthTemplate, TlsCertAuth, TokenAuth, TokenAuthToken, UniversalAuth, UniversalAuthClientSecret |
| `integration.infisical.upbound.io` | AwsParameterStore, AwsSecretsManager, Circleci, Databricks, GcpSecretManager |
| `kms.infisical.upbound.io` | ExternalAws, Key |
| `organization.infisical.upbound.io` | Alert, Role, SubOrganization, SubOrganizationGroup |
| `project.infisical.upbound.io` | AccessApprovalPolicy, Project, ProjectEnvironment, ProjectGroup, ProjectIdentity, ProjectIdentitySpecificPrivilege, ProjectRole, ProjectScopedIdentity, ProjectTemplate, ProjectUser, SecretApprovalPolicy, Webhook |
| `secret.infisical.upbound.io` | Secret, SecretFolder, SecretImport, SecretTag, SecretValidationRuleDynamicSecrets, SecretValidationRuleSecretRotations, SecretValidationRuleStaticSecrets |
| `secretrotation.infisical.upbound.io` | SecretRotationAuth0ClientSecret, SecretRotationAwsIamUserSecret, SecretRotationAzureClientSecret, SecretRotationDatadogApiKey, SecretRotationLdapPassword, SecretRotationLitellmApiKey, SecretRotationMssqlCredentials, SecretRotationMysqlCredentials, SecretRotationOracledbCredentials, SecretRotationPostgresCredentials |
| `secretsync.infisical.upbound.io` | SecretSyncAwsParameterStore, SecretSyncAwsSecretsManager, SecretSyncAzureAppConfiguration, SecretSyncAzureDevops, SecretSyncAzureKeyVault, SecretSyncBitbucket, SecretSyncCircleci, SecretSyncCloudflarePages, SecretSyncCloudflareWorkers, SecretSyncDatabricks, SecretSyncFlyio, SecretSyncGcpSecretManager, SecretSyncGithub, SecretSyncGitlab, SecretSyncOnePassword, SecretSyncRender, SecretSyncSpacelift, SecretSyncSupabase |

#### Namespaced managed resources

Every kind is also available as a namespaced managed resource in the
`*.infisical.m.upbound.io` API groups, for example
`project.infisical.m.upbound.io/v1beta1`. A namespaced managed resource
references a `ProviderConfig` in its own namespace, which reads the credentials
Secret from that namespace, or a cluster scoped `ClusterProviderConfig`:

```yaml
apiVersion: infisical.m.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: team-a
spec:
  host: https://app.infisical.com
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      key: credentials
---
apiVersion: project.infisical.m.upbound.io/v1beta1
kind: Project
metadata:
  name: example
  namespace: team-a
spec:
  forProvider:
    name: Example
    slug: example
    type: secret-manager
  providerConfigRef:
    kind: ProviderConfig
    name: default
```

See `examples/namespaced/providerconfig/` for a `ClusterProviderConfig`.

You can see the API reference [here](https://doc.crds.dev/github.com/infisical/provider-infisical).

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/infisical/provider-infisical/issues).
