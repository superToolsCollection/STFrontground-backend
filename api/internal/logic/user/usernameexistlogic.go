package logic

import (
	"STFrontground-backend/rpc/user/user"
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
	resp, err := l.svcCtx.User.UserNameExit(l.ctx, &user.UserNameExistReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserNameExistResp{
		IsExist: resp.IsExist,
	}, nil
}
