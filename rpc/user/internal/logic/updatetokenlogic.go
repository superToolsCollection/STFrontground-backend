package logic

import (
	"STFrontground-backend/rpc/user/model"
	"context"

	"STFrontground-backend/rpc/user/internal/svc"
	"STFrontground-backend/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(req *user.UpdateTokenReq) (*user.UpdateTokenResp, error) {
	userAuth := model.UsersAuth{
		UserId: req.UserId,
		Token:  req.Token,
	}
	err := l.svcCtx.UserAuthModel.Update(userAuth)
	if err != nil {
		return nil, err
	}
	return &user.UpdateTokenResp{
		UserId: req.UserId,
	}, nil
}
