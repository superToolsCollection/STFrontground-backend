package logic

import (
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetToolListByTagNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetToolListByTagNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetToolListByTagNameLogic {
	return &GetToolListByTagNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetToolListByTagNameLogic) GetToolListByTagName(in *tools.GetToolListByTagNameReq) (*tools.GetToolListByTagNameResp, error) {
	// todo: add your logic here and delete this line

	return &tools.GetToolListByTagNameResp{}, nil
}
