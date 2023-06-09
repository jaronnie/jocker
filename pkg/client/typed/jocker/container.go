// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package jocker

import (
	"context"

	"github.com/jaronnie/jocker/pkg/client/pb/jockerdpb"
	"github.com/jaronnie/jocker/pkg/client/rest"
)

type ContainerGetter interface {
	Container() ContainerInterface
}

type ContainerInterface interface {
	Container(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.Container, error)

	ContainerExpansion
}

type containerClient struct {
	client rest.Interface
}

func newContainerClient(c *JockerClient) *containerClient {
	return &containerClient{
		client: c.RESTClient(),
	}
}

func (x *containerClient) Container(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.Container, error) {
	var resp jockerdpb.Container
	err := x.client.Verb("POST").
		SubPath(
			"/api/v1/container",
		).
		Params().
		Body(param).
		Do(ctx).
		Into(&resp, false)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
