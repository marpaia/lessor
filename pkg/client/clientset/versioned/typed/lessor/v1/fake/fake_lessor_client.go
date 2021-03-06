package fake

import (
	v1 "github.com/lessor/lessor/pkg/client/clientset/versioned/typed/lessor/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLessorV1 struct {
	*testing.Fake
}

func (c *FakeLessorV1) Tenants(namespace string) v1.TenantInterface {
	return &FakeTenants{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLessorV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
