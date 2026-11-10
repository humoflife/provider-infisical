package dynamicsecret

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the dynamicsecret group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_dynamic_secret_aws_iam", func(r *config.Resource) {
		r.Kind = "AwsIam"
		r.ShortGroup = "dynamicsecret"
	})

	p.AddResourceConfigurator("infisical_dynamic_secret_kubernetes", func(r *config.Resource) {
		r.Kind = "Kubernetes"
		r.ShortGroup = "dynamicsecret"
	})

	p.AddResourceConfigurator("infisical_dynamic_secret_mongo_atlas", func(r *config.Resource) {
		r.Kind = "MongoAtlas"
		r.ShortGroup = "dynamicsecret"
	})

	p.AddResourceConfigurator("infisical_dynamic_secret_mongo_db", func(r *config.Resource) {
		r.Kind = "MongoDb"
		r.ShortGroup = "dynamicsecret"
	})

	p.AddResourceConfigurator("infisical_dynamic_secret_sql_database", func(r *config.Resource) {
		r.Kind = "SqlDatabase"
		r.ShortGroup = "dynamicsecret"
	})
}
