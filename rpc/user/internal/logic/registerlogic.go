package logic

import (
	"STFrontground-backend/rpc/pkg/errcode"
	"STFrontground-backend/rpc/user/model"
	"context"

	"STFrontground-backend/rpc/user/internal/svc"
	"STFrontground-backend/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(req *user.RegisterReq) (*user.RegisterResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	if err == nil {
		return nil, errcode.ErrorDuplicateUsername
	}

	_, err = l.svcCtx.UserModel.FindOneByMobile(req.Mobile)
	if err == nil {
		return nil, errcode.ErrorDuplicateMobile
	}
	_, err = l.svcCtx.UserModel.Insert(model.Users{
		Name:     req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, errcode.ErrorUserRegisterFail
	}
	return &user.RegisterResp{
		Isok: true,
	}, nil
}
