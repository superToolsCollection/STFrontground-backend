package logic

import (
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserNameExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserNameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserNameExistLogic {
	return UserNameExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserNameExistLogic) UserNameExist(req types.UserNameExistReq) (*types.UserNameExistResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserNameExistResp{}, nil
}
