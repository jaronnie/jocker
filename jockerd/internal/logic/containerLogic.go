package logic

import (
	"context"

	"github.com/jaronnie/jocker/jockerd/internal/svc"
	"github.com/jaronnie/jocker/jockerd/jockerdpb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerLogic {
	return &ContainerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContainerLogic) Container(in *jockerdpb.Empty) (*jockerdpb.Container, error) {
	// todo: add your logic here and delete this line

	return &jockerdpb.Container{}, nil
}
