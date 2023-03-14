package logic

import (
	"context"

	"github.com/jaronnie/jocker/jockerd/internal/svc"
	"github.com/jaronnie/jocker/jockerd/jockerd"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServerVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewServerVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServerVersionLogic {
	return &ServerVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ServerVersionLogic) ServerVersion(in *jockerd.Empty) (*jockerd.ServerVersionResponse, error) {
	// todo: add your logic here and delete this line

	return &jockerd.ServerVersionResponse{}, nil
}
