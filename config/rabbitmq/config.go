package rabbitmq

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure keeps every resource in the rabbitmq group and adds
// cross-resource references. Shared by the cluster and namespaced providers.
func Configure(p *config.Provider) {
	// upjet derives the group from the first word; keep every resource in the
	// rabbitmq group with a readable kind.
	for name, kind := range map[string]string{
		"rabbitmq_federation_upstream": "FederationUpstream",
		"rabbitmq_operator_policy":     "OperatorPolicy",
		"rabbitmq_topic_permissions":   "TopicPermissions",
	} {
		kind := kind
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.ShortGroup = "rabbitmq"
			r.Kind = kind
		})
	}
	for _, name := range []string{
		"rabbitmq_exchange", "rabbitmq_queue", "rabbitmq_policy", "rabbitmq_operator_policy",
		"rabbitmq_shovel", "rabbitmq_federation_upstream", "rabbitmq_binding",
	} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.References["vhost"] = config.Reference{TerraformName: "rabbitmq_vhost"}
		})
	}
	for _, name := range []string{"rabbitmq_permissions", "rabbitmq_topic_permissions"} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.References["vhost"] = config.Reference{TerraformName: "rabbitmq_vhost"}
			r.References["user"] = config.Reference{TerraformName: "rabbitmq_user"}
		})
	}
	p.AddResourceConfigurator("rabbitmq_topic_permissions", func(r *config.Resource) {
		r.References["permissions.exchange"] = config.Reference{TerraformName: "rabbitmq_exchange"}
	})
	p.AddResourceConfigurator("rabbitmq_binding", func(r *config.Resource) {
		r.References["source"] = config.Reference{TerraformName: "rabbitmq_exchange"}
		// destination may be a queue or an exchange, so no reference.
	})
}
