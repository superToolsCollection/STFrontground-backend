package logic

import (
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetToolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetToolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetToolListLogic {
	return &GetToolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetToolListLogic) GetToolList(in *tools.GetToolListReq) (*tools.GetToolListResp, error) {
	// todo: add your logic here and delete this line

	return &tools.GetToolListResp{}, nil
}
