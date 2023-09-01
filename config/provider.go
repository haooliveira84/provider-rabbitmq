/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/binding"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/exchange"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/federation_upstream"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/operator_policy"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/permissions"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/policy"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/queue"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/shovel"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/topic_permissions"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/user"
	"github.com/haooliveira84/provider-rabbitmq/config/rabbitmq/vhost"
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
		rabbitmq.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
