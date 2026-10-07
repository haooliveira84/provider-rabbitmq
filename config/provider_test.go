package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

var wantKinds = map[string]string{
	"rabbitmq_binding":             "Binding",
	tfExchange:                     "Exchange",
	"rabbitmq_federation_upstream": "FederationUpstream",
	"rabbitmq_operator_policy":     "OperatorPolicy",
	"rabbitmq_permissions":         "Permissions",
	"rabbitmq_policy":              "Policy",
	"rabbitmq_queue":               "Queue",
	"rabbitmq_shovel":              "Shovel",
	"rabbitmq_topic_permissions":   "TopicPermissions",
	tfUser:                         "User",
	tfVhost:                        "Vhost",
}

func TestProviders(t *testing.T) {
	for name, p := range map[string]*ujconfig.Provider{
		"cluster":    GetProvider(),
		"namespaced": GetProviderNamespaced(),
	} {
		t.Run(name, func(t *testing.T) {
			if len(p.Resources) != len(wantKinds) {
				t.Fatalf("generated %d resources, want %d", len(p.Resources), len(wantKinds))
			}
			for tfName, kind := range wantKinds {
				r, ok := p.Resources[tfName]
				if !ok {
					t.Errorf("%s: missing", tfName)
					continue
				}
				if r.Kind != kind {
					t.Errorf("%s: kind %q, want %q", tfName, r.Kind, kind)
				}
				if r.ShortGroup != "rabbitmq" {
					t.Errorf("%s: short group %q, want rabbitmq", tfName, r.ShortGroup)
				}
				if tfName != tfVhost && tfName != tfUser {
					if r.References["vhost"].TerraformName != tfVhost {
						t.Errorf("%s: vhost reference missing", tfName)
					}
				}
			}
			if p.Resources["rabbitmq_permissions"].References["user"].TerraformName != tfUser {
				t.Error("permissions: user reference missing")
			}
			if p.Resources["rabbitmq_topic_permissions"].References["permissions.exchange"].TerraformName != tfExchange {
				t.Error("topic_permissions: exchange reference missing")
			}
			if p.Resources["rabbitmq_binding"].References["source"].TerraformName != tfExchange {
				t.Error("binding: source reference missing")
			}
			if _, ok := p.Resources["rabbitmq_binding"].References["destination"]; ok {
				t.Error("binding: destination must not be a reference (queue or exchange)")
			}
		})
	}
}

func TestRootGroups(t *testing.T) {
	if g := GetProvider().RootGroup; g != "rabbitmq.crossplane.io" {
		t.Errorf("cluster root group %q", g)
	}
	if g := GetProviderNamespaced().RootGroup; g != "rabbitmq.m.crossplane.io" {
		t.Errorf("namespaced root group %q", g)
	}
}
