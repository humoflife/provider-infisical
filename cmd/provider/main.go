/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"os"
	"path/filepath"
	"time"

	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"gopkg.in/alecthomas/kingpin.v2"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	"terraform-provider-infisical/xpprovider"

	apisCluster "github.com/infisical/provider-infisical/apis/cluster"
	apisNamespaced "github.com/infisical/provider-infisical/apis/namespaced"
	"github.com/infisical/provider-infisical/config"
	"github.com/infisical/provider-infisical/internal/clients"
	controllerCluster "github.com/infisical/provider-infisical/internal/controller/cluster"
	controllerNamespaced "github.com/infisical/provider-infisical/internal/controller/namespaced"
	"github.com/infisical/provider-infisical/internal/features"
	"github.com/infisical/provider-infisical/internal/version"
)

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Infisical").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()
		metricsAddr             = app.Flag("metrics-bind-address", "The address the metric endpoint binds to.").Default("").Envar("METRICS_BIND_ADDRESS").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()

		healthProbeBindAddr = app.Flag("health-probe-bind-addr", "The address the health probe endpoints (/healthz, /readyz) listen on.").Default(":8081").Envar("HEALTH_PROBE_BIND_ADDRESS").String()

		// Every API has a single version, so the provider runs no conversion
		// webhook. These flags are kept as hidden no-ops, so that existing
		// deployments that still pass them do not fail to start.
		_ = app.Flag("webhook-port", "Deprecated: has no effect.").Hidden().Envar("WEBHOOK_PORT").Int()
		_ = app.Flag("certs-dir", "Deprecated: has no effect.").Hidden().Envar("CERTS_DIR").String()

		// External Secret Stores were removed in Crossplane v2. These flags
		// are kept as hidden no-ops, so that existing deployments that still
		// pass them do not fail to start.
		_                          = app.Flag("namespace", "Deprecated: has no effect.").Hidden().Default("crossplane-system").Envar("POD_NAMESPACE").String()
		enableExternalSecretStores = app.Flag("enable-external-secret-stores", "Deprecated: External Secret Stores are no longer supported, has no effect.").Hidden().Default("false").Envar("ENABLE_EXTERNAL_SECRET_STORES").Bool()
		_                          = app.Flag("ess-tls-cert-dir", "Deprecated: has no effect.").Hidden().Envar("ESS_TLS_CERTS_DIR").String()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	// The build sets the Terraform provider version from the Makefile, so
	// that it matches the version of the Terraform provider Go module.
	if version.TerraformProviderVersion == "" {
		kingpin.Fatalf("the Terraform provider version is not set: build the provider with make")
	}

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-infisical"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	controllerOpts := ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-infisical",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		HealthProbeBindAddress:     *healthProbeBindAddr,
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	}

	if metricsAddr != nil && *metricsAddr != "" {
		controllerOpts.Metrics = metricsserver.Options{
			BindAddress: *metricsAddr,
		}
	}

	mgr, err := ctrl.NewManager(cfg, controllerOpts)
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apisCluster.AddToScheme(mgr.GetScheme()), "Cannot add cluster scoped Infisical APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(mgr.GetScheme()), "Cannot add namespaced Infisical APIs to scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	// The Terraform provider runs in-process (no-fork). Its resources are
	// reconciled with the Terraform Plugin Framework, so no Terraform CLI
	// workspaces are needed.
	setupFn := clients.TerraformSetupBuilder(xpprovider.New(version.TerraformProviderVersion))
	operationTrackerStore := tjcontroller.NewOperationStore(log)
	newOptions := func(provider *ujconfig.Provider) tjcontroller.Options {
		return tjcontroller.Options{
			Options: xpcontroller.Options{
				Logger:                  log,
				GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
				PollInterval:            *pollInterval,
				MaxConcurrentReconciles: *maxReconcileRate,
				Features:                &feature.Flags{},
				MetricOptions: &xpcontroller.MetricOptions{
					PollStateMetricInterval: *pollStateMetricInterval,
					MRMetrics:               metricRecorder,
					MRStateMetrics:          stateMetrics,
				},
			},
			Provider:              provider,
			SetupFn:               setupFn,
			OperationTrackerStore: operationTrackerStore,
		}
	}
	clusterOpts := newOptions(config.GetProvider())
	namespacedOpts := newOptions(config.GetProviderNamespaced())

	if *enableExternalSecretStores {
		log.Info("External Secret Stores are no longer supported, ignoring the enable-external-secret-stores flag")
	}

	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOpts.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	kingpin.FatalIfError(mgr.AddReadyzCheck("ping", healthz.Ping), "Cannot add the readiness check")
	kingpin.FatalIfError(mgr.AddHealthzCheck("ping", healthz.Ping), "Cannot add the health check")

	kingpin.FatalIfError(controllerCluster.Setup(mgr, clusterOpts), "Cannot setup cluster scoped Infisical controllers")
	kingpin.FatalIfError(controllerNamespaced.Setup(mgr, namespacedOpts), "Cannot setup namespaced Infisical controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
