package svc

import (
	"github.com/fantopia-dev/website/internal/config"
	"github.com/fantopia-dev/website/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	TbWaitlistModel model.TbWaitlistModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// no cache, only database
	sqlConn := sqlx.NewMysql(c.MySql.DataSource)

	return &ServiceContext{
		Config: c,
		TbWaitlistModel: model.NewTbWaitlistModel(sqlConn),
	}
}
