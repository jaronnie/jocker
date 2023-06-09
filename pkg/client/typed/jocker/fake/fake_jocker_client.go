// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"github.com/jaronnie/jocker/pkg/client/rest"
	"github.com/jaronnie/jocker/pkg/client/typed/jocker"
)

type FakeJocker struct{}

func (f *FakeJocker) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (f *FakeJocker) Container() jocker.ContainerInterface {
	return &FakeContainer{Fake: f}
}

func (f *FakeJocker) Version() jocker.VersionInterface {
	return &FakeVersion{Fake: f}
}
