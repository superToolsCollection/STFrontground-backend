package logic

import (
	"STFrontground-backend/rpc/user/user"
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	resp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	jwtLogic := NewJwtLogic(l.ctx, l.svcCtx)
	r, err := jwtLogic.Jwt(types.JwtTokenRequest{})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Id:       resp.Id,
		Username: resp.Username,
		Mobile:   resp.Mobile,
		Token:    r.AccessToken,
	}, nil
}
