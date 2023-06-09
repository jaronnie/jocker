// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"context"

	"github.com/jaronnie/jocker/pkg/client/pb/jockerdpb"
)

var (
	FakeReturnServerVersion = &jockerdpb.ServerVersionResponse{}
)

type VersionGetter interface {
	Version() VersionInterface
}

type VersionInterface interface {
	ServerVersion(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.ServerVersionResponse, error)
}

type FakeVersion struct {
	Fake *FakeJocker
}

func (f *FakeVersion) ServerVersion(ctx context.Context, param *jockerdpb.Empty) (*jockerdpb.ServerVersionResponse, error) {
	return FakeReturnServerVersion, nil
}
