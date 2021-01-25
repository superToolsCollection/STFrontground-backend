package logic

import (
	"STFrontground-backend/rpc/user/user"
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(userId string) (*types.UserReply, error) {
	resp, err := l.svcCtx.User.UserInfo(l.ctx, &user.UserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserReply{
		Id:       resp.Id,
		Username: resp.Username,
		Mobile:   resp.Mobile,
	}, nil
}
