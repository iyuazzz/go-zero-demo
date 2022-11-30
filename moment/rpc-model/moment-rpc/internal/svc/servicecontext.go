package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	mysql_model "moment/mysql-model"
	"moment/rpc-model/moment-rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	MysqlModel mysql_model.MomentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MysqlModel: mysql_model.NewMomentModel(sqlx.NewMysql(c.DataSource)),
	}
}
