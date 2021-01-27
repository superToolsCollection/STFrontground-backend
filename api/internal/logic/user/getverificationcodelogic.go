package logic

import (
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetVerificationCodeLogic {
	return GetVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVerificationCodeLogic) GetVerificationCode(req types.GetVerificationCodeReq) (*types.GetVerificationCodeResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetVerificationCodeResp{}, nil
}
