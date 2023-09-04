package topicpermissions

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds topic_permissions configuration to the provider
func Configure(p *config.Provider) {
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
}
