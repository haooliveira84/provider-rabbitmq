package shovel

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds shovel configuration to the provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_shovel", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["vhost"] = config.Reference{
			Type: "VHost",
		}
		r.References["src_uri"] = config.Reference{
			Type: "SrcURI",
		}
		r.References["src_exchange"] = config.Reference{
			Type: "SrcExchange",
		}
		r.References["src_exchange_key"] = config.Reference{
			Type: "SrcExchangeKey",
		}
		r.References["dest_uri"] = config.Reference{
			Type: "DestURI",
		}
		r.References["dest_exchange"] = config.Reference{
			Type: "DestExchange",
		}
		r.References["dest_exchange_key"] = config.Reference{
			Type: "DestExchangeKey",
		}
		r.References["ack_mode"] = config.Reference{
			Type: "AckMode",
		}
		r.References["delete_after"] = config.Reference{
			Type: "DeleteAfter",
		}
		r.References["prefetch_count"] = config.Reference{
			Type: "PrefetchCount",
		}
		r.References["reconnect_delay"] = config.Reference{
			Type: "ReconnectDelay",
		}
		r.References["add_forward_headers"] = config.Reference{
			Type: "AddForwardHeaders",
		}
		r.References["ack_on_publish"] = config.Reference{
			Type: "AckOnPublish",
		}
		r.References["publish_properties"] = config.Reference{
			Type: "PublishProperties",
		}
		r.References["publish_fields"] = config.Reference{
			Type: "PublishFields",
		}
		r.References["ack_on_consume"] = config.Reference{
			Type: "AckOnConsume",
		}
		r.References["consume_properties"] = config.Reference{
			Type: "ConsumeProperties",
		}
		r.References["consume_fields"] = config.Reference{
			Type: "ConsumeFields",
		}
		r.References["reconnect_on_publish"] = config.Reference{
			Type: "ReconnectOnPublish",
		}
		r.References["reconnect_on_consume"] = config.Reference{
			Type: "Reconnect
