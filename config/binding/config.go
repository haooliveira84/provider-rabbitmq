package binding

import "github.com/upbound/upjet/pkg/config"

// Configure adds binding configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_binding", func(r *config.Resource) {
		r.References["source"] = config.Reference{
			Type: "Exchange",
		}
		r.References["destination"] = config.Reference{
			Type: "Exchange",
		}
		r.References["vhost"] = config.Reference{
			Type: "Vhost",
		}
		r.References["routing_key"] = config.Reference{
			Type: "RoutingKey",
		}
		r.References["destination_type"] = config.Reference{
			Type: "DestinationType",
		}
	})
}
