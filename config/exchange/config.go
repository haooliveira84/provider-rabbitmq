package exchange

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds exchange configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_exchange", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["vhost"] = config.Reference{
			Type: "VHost",
		}
		r.References["type"] = config.Reference{
			Type: "Type",
		}
		r.References["durable"] = config.Reference{
			Type: "Durable",
		}
		r.References["auto_delete"] = config.Reference{
			Type: "AutoDelete",
		}
		r.References["internal"] = config.Reference{
			Type: "Internal",
		}

	})
}
