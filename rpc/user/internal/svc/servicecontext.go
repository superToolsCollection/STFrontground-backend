package svc

import (
	"STFrontground-backend/rpc/user/internal/config"
	"STFrontground-backend/rpc/user/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UsersModel
	UserAuthModel model.UsersAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUsersModel(sqlx.NewMysql(c.DataSource)),
		UserAuthModel: model.NewUsersAuthModel(sqlx.NewMysql(c.DataSource)),
	}
}
