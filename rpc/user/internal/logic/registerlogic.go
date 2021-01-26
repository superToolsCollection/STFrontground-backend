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

	_, err = l.svcCtx.UserAuthModel.FindOneByPhone(req.Mobile)
	if err == nil {
		return nil, errcode.ErrorDuplicateMobile
	}

	result, err := l.svcCtx.UserModel.Insert(model.Users{
		Name:   req.Username,
		Gender: "保密",
		State:  "1",
	})
	if err != nil {
		return nil, errcode.ErrorUserRegisterFail
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errcode.ErrorUserRegisterFail
	}

	result, err = l.svcCtx.UserAuthModel.Insert(model.UsersAuth{
		UserId:        id,
		Name:          req.Username,
		Password:      req.Password,
		Phone:         req.Mobile,
		EmailActivate: "0",
	})
	if err != nil {
		return nil, errcode.ErrorUserRegisterFail
	}
	id, err = result.LastInsertId()
	if err != nil {
		return nil, errcode.ErrorUserRegisterFail
	}
	return &user.RegisterResp{
		Id:       id,
		Username: req.Username,
	}, nil
}
