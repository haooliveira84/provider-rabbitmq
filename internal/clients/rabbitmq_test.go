package clients

import (
	"context"
	"strings"
	"testing"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	clusterrmq "github.com/haooliveira84/provider-rabbitmq/apis/cluster/rabbitmq/v1alpha1"
	clusterv1beta1 "github.com/haooliveira84/provider-rabbitmq/apis/cluster/v1beta1"
	nsrmq "github.com/haooliveira84/provider-rabbitmq/apis/namespaced/rabbitmq/v1alpha1"
	namespacedv1beta1 "github.com/haooliveira84/provider-rabbitmq/apis/namespaced/v1beta1"
)

const creds = `{"endpoint":"http://rmq:15672","username":"admin","password":"s3cret","insecure":"true"}`

func newClient(t *testing.T, objs ...client.Object) client.Client {
	t.Helper()
	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{
		corev1.AddToScheme, clusterv1beta1.SchemeBuilder.AddToScheme, namespacedv1beta1.SchemeBuilder.AddToScheme,
		clusterrmq.SchemeBuilder.AddToScheme, nsrmq.SchemeBuilder.AddToScheme,
	} {
		if err := add(s); err != nil {
			t.Fatal(err)
		}
	}
	return fake.NewClientBuilder().WithScheme(s).WithObjects(objs...).Build()
}

func secret(ns string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: "creds", Namespace: ns},
		Data:       map[string][]byte{"credentials": []byte(creds)},
	}
}

func assertConfig(t *testing.T, got map[string]any) {
	t.Helper()
	want := map[string]any{"endpoint": "http://rmq:15672", "username": "admin", "password": "s3cret", "insecure": "true"}
	if len(got) != len(want) {
		t.Fatalf("configuration = %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %v, want %v", k, got[k], v)
		}
	}
}

func TestSetupClusterScoped(t *testing.T) {
	pc := &clusterv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default"},
		Spec: clusterv1beta1.ProviderConfigSpec{Credentials: clusterv1beta1.ProviderCredentials{
			Source: xpv2.CredentialsSourceSecret,
			CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
				SecretReference: xpv2.SecretReference{Name: "creds", Namespace: "crossplane-system"}, Key: "credentials",
			}},
		}},
	}
	mg := &clusterrmq.Vhost{ObjectMeta: metav1.ObjectMeta{Name: "prod", UID: "uid-prod"}}
	mg.Spec.ProviderConfigReference = &xpv2.Reference{Name: "default"}

	ps, err := TerraformSetupBuilder("1.5.7", "cyrilgdn/rabbitmq", "1.10.1")(context.Background(), newClient(t, pc, secret("crossplane-system"), mg), mg)
	if err != nil {
		t.Fatal(err)
	}
	if ps.Requirement.Source != "cyrilgdn/rabbitmq" || ps.Requirement.Version != "1.10.1" || ps.Version != "1.5.7" {
		t.Errorf("requirement = %+v", ps.Requirement)
	}
	assertConfig(t, ps.Configuration)
}

func TestSetupNamespaced(t *testing.T) {
	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "team-a"},
		Spec: namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
			Source: xpv2.CredentialsSourceSecret,
			CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
				// Namespace is ignored: the MR namespace wins.
				SecretReference: xpv2.SecretReference{Name: "creds", Namespace: "elsewhere"}, Key: "credentials",
			}},
		}},
	}
	mg := &nsrmq.Vhost{ObjectMeta: metav1.ObjectMeta{Name: "prod", Namespace: "team-a", UID: "uid-prod"}}
	mg.Spec.ProviderConfigReference = &xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: "default"}

	ps, err := TerraformSetupBuilder("1.5.7", "cyrilgdn/rabbitmq", "1.10.1")(context.Background(), newClient(t, pc, secret("team-a"), mg), mg)
	if err != nil {
		t.Fatal(err)
	}
	assertConfig(t, ps.Configuration)
}

func TestSetupMissingProviderConfigRef(t *testing.T) {
	mg := &clusterrmq.Vhost{ObjectMeta: metav1.ObjectMeta{Name: "prod", UID: "uid-prod"}}
	_, err := TerraformSetupBuilder("", "", "")(context.Background(), newClient(t, mg), mg)
	if err == nil || !strings.Contains(err.Error(), errNoProviderConfig) {
		t.Fatalf("err = %v, want %q", err, errNoProviderConfig)
	}
}
