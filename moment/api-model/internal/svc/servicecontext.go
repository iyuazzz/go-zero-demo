package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"moment/api-model/internal/config"
	"moment/rpc-model/moment-rpc/momentclient"
	"moment/rpc-model/token-rpc/tokenclient"
)

type ServiceContext struct {
	Config    config.Config
	MomentRpc momentclient.Moment
	TokenRpc  tokenclient.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MomentRpc: momentclient.NewMoment(zrpc.MustNewClient(c.MomentRpc)),
		TokenRpc:  tokenclient.NewToken(zrpc.MustNewClient(c.TokenRpc)),
	}
}
