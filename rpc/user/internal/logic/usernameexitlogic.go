package logic

import (
	"context"

	"STFrontground-backend/rpc/user/internal/svc"
	"STFrontground-backend/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserNameExitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserNameExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserNameExitLogic {
	return &UserNameExitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserNameExitLogic) UserNameExit(req *user.UserNameExistReq) (*user.UserNameExistResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	if err != nil {
		return nil, err
	}
	isExist := false
	if userInfo != nil {
		isExist = true
	}
	return &user.UserNameExistResp{
		IsExist: isExist,
	}, nil
}
