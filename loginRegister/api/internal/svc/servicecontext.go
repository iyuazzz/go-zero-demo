package svc

import (
	"loginRegister/api/internal/config"
	"loginRegister/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel("mongodb://admin:123456@0.0.0.0:27017", "loginAndRegister", "user"),
	}
}
