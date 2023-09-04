package permissions

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds permissions configuration to the provider
func Configure(p *config.Provider) {
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
}
