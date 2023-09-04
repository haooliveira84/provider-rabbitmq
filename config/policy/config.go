package policy

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds policy configuration to the provider
func Configure(p *config.Provider) {
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
}
