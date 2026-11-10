package gateway

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the resources of the gateway group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_gateway", func(r *config.Resource) {
		r.Kind = "Gateway"
		r.ShortGroup = "gateway"
	})

	p.AddResourceConfigurator("infisical_gateway_enrollment_token", func(r *config.Resource) {
		r.Kind = "EnrollmentToken"
		r.ShortGroup = "gateway"
		r.References["gateway_id"] = config.Reference{TerraformName: "infisical_gateway"}
	})
}
