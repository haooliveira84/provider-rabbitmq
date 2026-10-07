package config

import (
	"context"
	"testing"
)

const (
	vhost    = "prod"
	vhostKey = "vhost"
)

func TestExternalNames(t *testing.T) {
	cases := map[string]struct {
		resource     string
		externalName string
		params       map[string]any
		wantID       string
	}{
		"vhost":              {tfVhost, vhost, nil, vhost},
		"user":               {tfUser, "alice", nil, "alice"},
		"exchange":           {tfExchange, "orders", map[string]any{vhostKey: vhost}, "orders@prod"},
		"queue":              {"rabbitmq_queue", "orders.created", map[string]any{vhostKey: vhost}, "orders.created@prod"},
		"policy":             {"rabbitmq_policy", "ha", map[string]any{"vhost": "/"}, "ha@/"},
		"operatorPolicy":     {"rabbitmq_operator_policy", "limits", map[string]any{vhostKey: vhost}, "limits@prod"},
		"shovel":             {"rabbitmq_shovel", "mover", map[string]any{vhostKey: vhost}, "mover@prod"},
		"federationUpstream": {"rabbitmq_federation_upstream", "up", map[string]any{vhostKey: vhost}, "up@prod"},
		"permissions":        {"rabbitmq_permissions", "alice", map[string]any{vhostKey: vhost}, "alice@prod"},
		"topicPermissions":   {"rabbitmq_topic_permissions", "alice", map[string]any{vhostKey: vhost}, "alice@prod"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			en, ok := ExternalNameConfigs[tc.resource]
			if !ok {
				t.Fatalf("no external name config for %s", tc.resource)
			}
			id, err := en.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			if err != nil {
				t.Fatalf("GetIDFn: %v", err)
			}
			if id != tc.wantID {
				t.Fatalf("GetIDFn = %q, want %q", id, tc.wantID)
			}
			got, err := en.GetExternalNameFn(map[string]any{"id": id})
			if err != nil {
				t.Fatalf("GetExternalNameFn: %v", err)
			}
			if got != tc.externalName {
				t.Fatalf("GetExternalNameFn = %q, want %q", got, tc.externalName)
			}
		})
	}
}

func TestBindingUsesProviderID(t *testing.T) {
	en := ExternalNameConfigs["rabbitmq_binding"]
	const id = "prod/orders/orders.created/queue/%23"
	got, err := en.GetExternalNameFn(map[string]any{"id": id})
	if err != nil || got != id {
		t.Fatalf("GetExternalNameFn = %q, %v; want %q", got, err, id)
	}
}

func TestExternalNameConfiguredCoversSchema(t *testing.T) {
	want := 11
	if got := len(ExternalNameConfigured()); got != want {
		t.Fatalf("configured %d resources, want %d", got, want)
	}
}
