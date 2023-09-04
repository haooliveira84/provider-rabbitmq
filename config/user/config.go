package user

import "github.com/upbound/upjet/pkg/config"

// Configure adds user configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_user", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.Sensitive.AddFieldPath("password")
		r.References["password"] = config.Reference{
			Type: "Password",
		}
		r.References["tags"] = config.Reference{
			Type: "Tags",
		}
	})
}
