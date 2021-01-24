package logic

import (
	"STFrontground-backend/rpc/pkg/errcode"
	"STFrontground-backend/rpc/user/model"
	"context"

	"STFrontground-backend/rpc/user/internal/svc"
	"STFrontground-backend/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(req *user.LoginReq) (*user.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	switch err {
	case nil:
		if userInfo.Password != req.Password {
			return nil, errcode.ErrorUserIncorrectPassword
		}
		return &user.LoginResp{
			Id:       userInfo.Id,
			Username: userInfo.Name,
			Mobile:   userInfo.Mobile,
		}, nil
	case model.ErrNotFound:
		return nil, errcode.ErrorUsernameUnRegister
	default:
		return nil, err
	}
}
