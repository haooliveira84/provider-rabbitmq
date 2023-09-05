package rabbitmq

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
	p.AddResourceConfigurator("rabbitmq_operator_policy", func(r *config.Resource) {
		r.References = config.References{
			"name": {
				Type: "Name",
			},
			"vhost": {
				Type: "VHost",
			},
			"policy.pattern": {
				Type: "PolicyPattern",
			},
			"policy.priority": {
				Type: "PolicyPriority",
			},
			"policy.definition.message-ttl": {
				Type: "PolicyDefinitionMessageTTL",
			},
			"policy.definition.max-length": {
				Type: "PolicyDefinitionMaxLength",
			},
			"policy.definition.expires": {
				Type: "PolicyDefinitionExpires",
			},
		}
	})
	p.AddResourceConfigurator("rabbitmq_permissions", func(r *config.Resource) {
		r.References = config.References{
			"user": {
				Type: "User",
			},
			"vhost": {
				Type: "VHost",
			},
			"permissions.configure": {
				Type: "PermissionsConfigure",
			},
			"permissions.write": {
				Type: "PermissionsWrite",
			},
			"permissions.read": {
				Type: "PermissionsRead",
			},
		}
	})
	p.AddResourceConfigurator("rabbitmq_policy", func(r *config.Resource) {
		r.References = config.References{
			"name": {
				Type: "Name",
			},
			"vhost": {
				Type: "VHost",
			},
			"policy.pattern": {
				Type: "PolicyPattern",
			},
			"policy.priority": {
				Type: "PolicyPriority",
			},
			"policy.definition.message-ttl": {
				Type: "PolicyDefinitionMessageTTL",
			},
			"policy.definition.max-length": {
				Type: "PolicyDefinitionMaxLength",
			},
			"policy.definition.expires": {
				Type: "PolicyDefinitionExpires",
			},
			"policy.definition.dead-letter-exchange": {
				Type: "PolicyDefinitionDeadLetterExchange",
			},
			"policy.definition.dead-letter-routing-key": {
				Type: "PolicyDefinitionDeadLetterRoutingKey",
			},
			"policy.definition.ha-mode": {
				Type: "PolicyDefinitionHAMode",
			},
			"policy.definition.ha-params": {
				Type: "PolicyDefinitionHAParams",
			},
			"policy.definition.federation-upstream-set": {
				Type: "PolicyDefinitionFederationUpstreamSet",
			},
			"policy.definition.federation-upstream": {
				Type: "PolicyDefinitionFederationUpstream",
			},
		}
	})
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
			Type: "Reconnect",
		}
	})
	p.AddResourceConfigurator("rabbitmq_topic_permissions", func(r *config.Resource) {
		r.References = config.References{
			"user": {
				Type: "User",
			},
			"vhost": {
				Type: "VHost",
			},
			"exchange": {
				Type: "Exchange",
			},
			"write": {
				Type: "Write",
			},
			"read": {
				Type: "Read",
			},
			"configure": {
				Type: "Configure",
			},
		}
	})
	p.AddResourceConfigurator("rabbitmq_user", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["password"] = config.Reference{
			Type: "Password",
		}
		r.References["tags"] = config.Reference{
			Type: "Tags",
		}
	})
	p.AddResourceConfigurator("rabbitmq_vhost", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "Name",
		}
	})
}
