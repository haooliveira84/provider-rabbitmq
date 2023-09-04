package federationupstream

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds federation upstream configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_federation_upstream", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["vhost"] = config.Reference{
			Type: "VHost",
		}
		r.References["uri"] = config.Reference{
			Type: "URI",
		}
		r.References["prefetch_count"] = config.Reference{
			Type: "PrefetchCount",
		}
		r.References["reconnect_delay"] = config.Reference{
			Type: "ReconnectDelay",
		}
		r.References["ack_mode"] = config.Reference{
			Type: "AckMode",
		}
		r.References["trust_user_id"] = config.Reference{
			Type: "TrustUserID",
		}
		r.References["max_hops"] = config.Reference{
			Type: "MaxHops",
		}
		r.References["expires"] = config.Reference{
			Type: "Expires",
		}
		r.References["message_ttl"] = config.Reference{
			Type: "MessageTTL",
		}
	})
}
