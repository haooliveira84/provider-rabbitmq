// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/providerconfig"
	binding "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/binding"
	exchange "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/exchange"
	federationupstream "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/federationupstream"
	operatorpolicy "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/operatorpolicy"
	permissions "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/permissions"
	policy "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/policy"
	queue "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/queue"
	shovel "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/shovel"
	topicpermissions "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/topicpermissions"
	user "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/user"
	vhost "github.com/haooliveira84/provider-rabbitmq/internal/controller/cluster/rabbitmq/vhost"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		binding.Setup,
		exchange.Setup,
		federationupstream.Setup,
		operatorpolicy.Setup,
		permissions.Setup,
		policy.Setup,
		queue.Setup,
		shovel.Setup,
		topicpermissions.Setup,
		user.Setup,
		vhost.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		binding.SetupGated,
		exchange.SetupGated,
		federationupstream.SetupGated,
		operatorpolicy.SetupGated,
		permissions.SetupGated,
		policy.SetupGated,
		queue.SetupGated,
		shovel.SetupGated,
		topicpermissions.SetupGated,
		user.SetupGated,
		vhost.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		providerconfig.SetupWebhookWithManager,
		binding.SetupWebhookWithManager,
		exchange.SetupWebhookWithManager,
		federationupstream.SetupWebhookWithManager,
		operatorpolicy.SetupWebhookWithManager,
		permissions.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		queue.SetupWebhookWithManager,
		shovel.SetupWebhookWithManager,
		topicpermissions.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		vhost.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
