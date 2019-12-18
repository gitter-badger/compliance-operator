// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/compliance-operator/pkg/apis/machineconfiguration/v1"
	"github.com/openshift/compliance-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type MachineconfigurationV1Interface interface {
	RESTClient() rest.Interface
	MachineConfigsGetter
	MachineConfigPoolsGetter
}

// MachineconfigurationV1Client is used to interact with features provided by the machineconfiguration.openshift.io group.
type MachineconfigurationV1Client struct {
	restClient rest.Interface
}

func (c *MachineconfigurationV1Client) MachineConfigs() MachineConfigInterface {
	return newMachineConfigs(c)
}

func (c *MachineconfigurationV1Client) MachineConfigPools() MachineConfigPoolInterface {
	return newMachineConfigPools(c)
}

// NewForConfig creates a new MachineconfigurationV1Client for the given config.
func NewForConfig(c *rest.Config) (*MachineconfigurationV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MachineconfigurationV1Client{client}, nil
}

// NewForConfigOrDie creates a new MachineconfigurationV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MachineconfigurationV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MachineconfigurationV1Client for the given RESTClient.
func New(c rest.Interface) *MachineconfigurationV1Client {
	return &MachineconfigurationV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MachineconfigurationV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
