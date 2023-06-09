// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package jocker

import (
	"context"

	"github.com/jaronnie/jocker/pkg/client/pb/jockerdpb"
	"github.com/jaronnie/jocker/pkg/client/rest"
)

type VersionGetter interface {
	Version() VersionInterface
}

type VersionInterface interface {
	ServerVersion(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.ServerVersionResponse, error)

	VersionExpansion
}

type versionClient struct {
	client rest.Interface
}

func newVersionClient(c *JockerClient) *versionClient {
	return &versionClient{
		client: c.RESTClient(),
	}
}

func (x *versionClient) ServerVersion(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.ServerVersionResponse, error) {
	var resp jockerdpb.ServerVersionResponse
	err := x.client.Verb("POST").
		SubPath(
			"/api/v1/version",
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
