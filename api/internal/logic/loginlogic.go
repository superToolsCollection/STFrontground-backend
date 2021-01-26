package logic

import (
	"STFrontground-backend/rpc/pkg/errcode"
	"STFrontground-backend/rpc/user/user"
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

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
		return nil, errcode.ErrorUserIncorrectPassword
	}
	//获取jwt-token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, resp.Id)
	if err != nil {
		return nil, errcode.UnauthorizedTokenGenerate
	}
	return &types.LoginResp{
		Id:       resp.Id,
		Username: resp.Username,
		JwtToken: types.JwtToken{
			AccessToken: jwtToken,
		},
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
