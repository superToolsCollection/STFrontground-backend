package logic

import (
	"context"
	"strconv"

	"STFrontground-backend/rpc/user/internal/svc"
	"STFrontground-backend/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(req *user.UserInfoReq) (*user.UserInfoResp, error) {
	userId, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(userId)
	if err != nil {
		return nil, err
	}
	return &user.UserInfoResp{
		Id:       userInfo.Id,
		Username: userInfo.Name,
		Mobile:   userInfo.Mobile,
	}, nil
}
