package logic

import (
	"context"

	"github.com/jaronnie/jocker/jockerd/internal/svc"
	"github.com/jaronnie/jocker/jockerd/jockerdpb"

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

func (l *ServerVersionLogic) ServerVersion(in *jockerdpb.Empty) (*jockerdpb.ServerVersionResponse, error) {
	return &jockerdpb.ServerVersionResponse{Version: "beta testing!"}, nil
}
