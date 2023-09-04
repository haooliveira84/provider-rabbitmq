/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/haooliveira84/provider-rabbitmq/config/binding"
	"github.com/haooliveira84/provider-rabbitmq/config/exchange"
	federationupstream "github.com/haooliveira84/provider-rabbitmq/config/federation_upstream"
	operatorpolicy "github.com/haooliveira84/provider-rabbitmq/config/operator_policy"
	"github.com/haooliveira84/provider-rabbitmq/config/permissions"
	"github.com/haooliveira84/provider-rabbitmq/config/policy"
	"github.com/haooliveira84/provider-rabbitmq/config/queue"
	"github.com/haooliveira84/provider-rabbitmq/config/shovel"
	topicpermissions "github.com/haooliveira84/provider-rabbitmq/config/topic_permissions"
	"github.com/haooliveira84/provider-rabbitmq/config/user"
	"github.com/haooliveira84/provider-rabbitmq/config/vhost"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "template"
	modulePath     = "github.com/haooliveira84/provider-rabbitmq"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		binding.Configure,
		exchange.Configure,
		federationupstream.Configure,
		operatorpolicy.Configure,
		permissions.Configure,
		policy.Configure,
		queue.Configure,
		shovel.Configure,
		topicpermissions.Configure,
		user.Configure,
		vhost.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
