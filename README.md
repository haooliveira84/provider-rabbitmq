# provider-rabbitmq

A [Crossplane](https://crossplane.io/) provider for RabbitMQ, generated with
[Upjet](https://github.com/crossplane/upjet) from
[terraform-provider-rabbitmq](https://github.com/cyrilgdn/terraform-provider-rabbitmq).
It talks to the RabbitMQ management HTTP API and manages vhosts, users,
permissions, exchanges, queues, bindings, policies, shovels and federation
upstreams as Kubernetes resources.

## Managed resources

Every kind is available in two API groups:

| Scope | API group | ProviderConfig kinds |
| --- | --- | --- |
| Cluster-scoped (legacy) | `rabbitmq.rabbitmq.crossplane.io/v1alpha1` | `ProviderConfig` in `rabbitmq.crossplane.io/v1beta1` |
| Namespaced | `rabbitmq.rabbitmq.m.crossplane.io/v1alpha1` | `ProviderConfig` or `ClusterProviderConfig` in `rabbitmq.m.crossplane.io/v1beta1` |

| Kind | Terraform resource | External name |
| --- | --- | --- |
| `Vhost` | `rabbitmq_vhost` | `<name>` |
| `User` | `rabbitmq_user` | `<name>` |
| `Permissions` | `rabbitmq_permissions` | `<user>@<vhost>` |
| `TopicPermissions` | `rabbitmq_topic_permissions` | `<user>@<vhost>` |
| `Exchange` | `rabbitmq_exchange` | `<name>@<vhost>` |
| `Queue` | `rabbitmq_queue` | `<name>@<vhost>` |
| `Binding` | `rabbitmq_binding` | assigned by RabbitMQ (`vhost/source/destination/type/properties_key`) |
| `Policy` | `rabbitmq_policy` | `<name>@<vhost>` |
| `OperatorPolicy` | `rabbitmq_operator_policy` | `<name>@<vhost>` |
| `Shovel` | `rabbitmq_shovel` | `<name>@<vhost>` |
| `FederationUpstream` | `rabbitmq_federation_upstream` | `<name>@<vhost>` |

The external name is `metadata.name` unless you set the
`crossplane.io/external-name` annotation. For the `@<vhost>` kinds the
`vhost` part comes from `spec.forProvider.vhost`, so the name is only the
RabbitMQ object name. Cross-resource references are available as
`vhostRef`/`vhostSelector`, `userRef`/`userSelector` (permissions) and
`exchangeRef`/`exchangeSelector` (topic permissions, binding source).

Example manifests for every kind live in
[examples-generated/](examples-generated/).

## Installation

```sh
kubectl apply -f - <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-rabbitmq
spec:
  package: haooliveira84/provider-rabbitmq:v0.1.0
EOF
```

## Configuration

Credentials are a JSON document in a Secret. All keys accepted by the
Terraform provider are passed through:

| Key | Required | Description |
| --- | --- | --- |
| `endpoint` | yes | Management API URL, e.g. `http://rabbitmq.default.svc:15672` |
| `username` | yes | Management user |
| `password` | yes | Management password |
| `insecure` | no | `"true"` to skip TLS verification |
| `cacert_file` | no | Path to a CA bundle inside the provider pod |
| `clientcert_file` | no | Path to a client certificate |
| `clientkey_file` | no | Path to a client key |
| `proxy` | no | HTTP proxy URL |

```sh
kubectl -n crossplane-system create secret generic rabbitmq-creds \
  --from-literal=credentials='{"endpoint":"http://rabbitmq.default.svc:15672","username":"admin","password":"s3cret"}'
```

Cluster-scoped `ProviderConfig`:

```yaml
apiVersion: rabbitmq.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: rabbitmq-creds
      namespace: crossplane-system
      key: credentials
```

Namespaced `ProviderConfig` (the Secret must be in the same namespace as the
managed resources):

```yaml
apiVersion: rabbitmq.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: team-a
spec:
  credentials:
    source: Secret
    secretRef:
      name: rabbitmq-creds
      key: credentials
```

## Usage

```yaml
apiVersion: rabbitmq.rabbitmq.crossplane.io/v1alpha1
kind: Vhost
metadata:
  name: orders
spec:
  forProvider: {}
---
apiVersion: rabbitmq.rabbitmq.crossplane.io/v1alpha1
kind: User
metadata:
  name: orders-app
spec:
  forProvider:
    passwordSecretRef:
      name: orders-app-password
      namespace: crossplane-system
      key: password
    tags: [management]
---
apiVersion: rabbitmq.rabbitmq.crossplane.io/v1alpha1
kind: Permissions
metadata:
  name: orders-app
spec:
  forProvider:
    userRef: {name: orders-app}
    vhostRef: {name: orders}
    permissions:
      - configure: ".*"
        write: ".*"
        read: ".*"
---
apiVersion: rabbitmq.rabbitmq.crossplane.io/v1alpha1
kind: Queue
metadata:
  name: orders.created
spec:
  forProvider:
    vhostRef: {name: orders}
    settings:
      - durable: true
        arguments:
          x-queue-type: quorum
```

## Development

Requirements: Go (see `go.mod`), Docker, `make`. The build submodule is
fetched on the first `make` run.

| Command | What it does |
| --- | --- |
| `make generate` | Fetches the Terraform provider schema and docs, regenerates APIs, controllers, CRDs and examples |
| `make build` | Builds the provider binary and image |
| `make test` / `go test ./...` | Unit tests for external names, resource configuration and credential handling |
| `make lint` | golangci-lint |
| `make run` | Runs the provider out-of-cluster against the current kubeconfig |
| `make e2e` | Spins up a kind cluster and runs uptest against `examples/` |

Resource configuration lives in [config/](config/):

- [external_name.go](config/external_name.go): how Terraform IDs map to
  Crossplane external names.
- [rabbitmq/config.go](config/rabbitmq/config.go): kinds, API group and
  cross-resource references.
- [provider.go](config/provider.go): wires the two providers (cluster and
  namespaced).

Upstream versions (Terraform CLI and provider) are pinned in the
[Makefile](Makefile). Bump `TERRAFORM_PROVIDER_VERSION` and
`TERRAFORM_NATIVE_PROVIDER_BINARY` together, then run `make generate`.

## Report a bug

Open an [issue](https://github.com/haooliveira84/provider-rabbitmq/issues).
