package operatorpolicy

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds operator policy configuration to the provider
func Configure(p *config.Provider) {
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
}
