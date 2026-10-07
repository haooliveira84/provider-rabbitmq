package config

import "github.com/crossplane/upjet/v2/pkg/config"

// Terraform resource names that are referenced from more than one place.
const (
	tfVhost    = "rabbitmq_vhost"
	tfUser     = "rabbitmq_user"
	tfExchange = "rabbitmq_exchange"
)

// nameAtVhost: Terraform ID is "<name>@<vhost>".
var nameAtVhost = config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}")

// userAtVhost: Terraform ID is "<user>@<vhost>".
var userAtVhost = config.TemplatedStringAsIdentifier("user", "{{ .external_name }}@{{ .parameters.vhost }}")

// ExternalNameConfigs contains all external name configurations for this
// provider. IDs follow the import formats documented by
// terraform-provider-rabbitmq.
var ExternalNameConfigs = map[string]config.ExternalName{
	tfVhost: config.NameAsIdentifier,
	tfUser:  config.NameAsIdentifier,

	tfExchange:                     nameAtVhost,
	"rabbitmq_queue":               nameAtVhost,
	"rabbitmq_policy":              nameAtVhost,
	"rabbitmq_operator_policy":     nameAtVhost,
	"rabbitmq_shovel":              nameAtVhost,
	"rabbitmq_federation_upstream": nameAtVhost,

	"rabbitmq_permissions":       userAtVhost,
	"rabbitmq_topic_permissions": userAtVhost,

	// ID is "vhost/source/destination/destination_type/properties_key" and
	// properties_key is computed by the server.
	"rabbitmq_binding": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
