package svc

import (
	"STFrontground-backend/rpc/user/internal/config"
	"STFrontground-backend/rpc/user/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
