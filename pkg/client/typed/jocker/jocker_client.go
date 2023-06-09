// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package jocker

import (
	"github.com/jaronnie/jocker/pkg/client/rest"
)

type JockerInterface interface {
	RESTClient() rest.Interface

	ContainerGetter
	VersionGetter
}

type JockerClient struct {
	restClient rest.Interface
}

func (x *JockerClient) RESTClient() rest.Interface {
	if x == nil {
		return nil
	}
	return x.restClient
}

func (x *JockerClient) Container() ContainerInterface {
	return newContainerClient(x)
}

func (x *JockerClient) Version() VersionInterface {
	return newVersionClient(x)
}

// NewForConfig creates a new JockerClient for the given config.
func NewForConfig(x *rest.RESTClient) (*JockerClient, error) {
	config := *x
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &JockerClient{client}, nil
}
