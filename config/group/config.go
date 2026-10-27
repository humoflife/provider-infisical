package group

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_group", func(r *config.Resource) {
		r.Kind = "Group"
		r.ShortGroup = "group"
	})

	p.AddResourceConfigurator("infisical_group_machine_identity", func(r *config.Resource) {
		r.Kind = "GroupMachineIdentity"
		r.ShortGroup = "group"
		r.References["group_id"] = config.Reference{TerraformName: "infisical_group"}
		r.References["identity_id"] = config.Reference{TerraformName: "infisical_identity"}
	})
}
