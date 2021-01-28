package svc

import (
	"STFrontground-backend/rpc/tools/internal/config"
	"STFrontground-backend/rpc/tools/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	StoriesModel model.StoriesModel
	ForbesModel  model.ForbesModel
	TagsModel    model.TagsModel
	ToolsModel   model.ToolsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		StoriesModel: model.NewStoriesModel(sqlx.NewMysql(c.DataSource)),
		ForbesModel:  model.NewForbesModel(sqlx.NewMysql(c.DataSource)),
		TagsModel:    model.NewTagsModel(sqlx.NewMysql(c.DataSource)),
		ToolsModel:   model.NewToolsModel(sqlx.NewMysql(c.DataSource)),
	}
}
