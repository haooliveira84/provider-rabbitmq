package vhost

import "github.com/upbound/upjet/pkg/config"

// Configure adds vhost configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_vhost", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
	})
}
