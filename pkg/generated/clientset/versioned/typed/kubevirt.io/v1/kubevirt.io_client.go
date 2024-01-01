/*
Copyright 2024 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"net/http"

	"github.com/starbops/vm-dhcp-controller/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
	v1 "kubevirt.io/api/core/v1"
)

type KubevirtV1Interface interface {
	RESTClient() rest.Interface
	KubeVirtsGetter
	VirtualMachinesGetter
	VirtualMachineInstancesGetter
	VirtualMachineInstanceMigrationsGetter
	VirtualMachineInstancePresetsGetter
	VirtualMachineInstanceReplicaSetsGetter
}

// KubevirtV1Client is used to interact with features provided by the kubevirt.io group.
type KubevirtV1Client struct {
	restClient rest.Interface
}

func (c *KubevirtV1Client) KubeVirts(namespace string) KubeVirtInterface {
	return newKubeVirts(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachines(namespace string) VirtualMachineInterface {
	return newVirtualMachines(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstances(namespace string) VirtualMachineInstanceInterface {
	return newVirtualMachineInstances(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstanceMigrations(namespace string) VirtualMachineInstanceMigrationInterface {
	return newVirtualMachineInstanceMigrations(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstancePresets(namespace string) VirtualMachineInstancePresetInterface {
	return newVirtualMachineInstancePresets(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetInterface {
	return newVirtualMachineInstanceReplicaSets(c, namespace)
}

// NewForConfig creates a new KubevirtV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KubevirtV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KubevirtV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KubevirtV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KubevirtV1Client{client}, nil
}

// NewForConfigOrDie creates a new KubevirtV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubevirtV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubevirtV1Client for the given RESTClient.
func New(c rest.Interface) *KubevirtV1Client {
	return &KubevirtV1Client{c}
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
func (c *KubevirtV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
