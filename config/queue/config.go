package queue

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds queue configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_queue", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["vhost"] = config.Reference{
			Type: "VHost",
		}
		r.References["durable"] = config.Reference{
			Type: "Durable",
		}
		r.References["auto_delete"] = config.Reference{
			Type: "AutoDelete",
		}
		r.References["exclusive"] = config.Reference{
			Type: "Exclusive",
		}
		r.References["arguments"] = config.Reference{
			Type: "Arguments",
		}
	})
}
