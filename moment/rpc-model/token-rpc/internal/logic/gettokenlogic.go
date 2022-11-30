package logic

import (
	"context"
	"fmt"

	"moment/rpc-model/token-rpc/internal/svc"
	"moment/rpc-model/token-rpc/types/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *token.TokenRequest) (*token.TokenReply, error) {
	ret := &token.TokenReply{
		Token: "iyua",
	}
	if in.Id != "iyua" {
		ret.Token = ""
		ret.Result = fmt.Sprintf("fail to get token, user(%s) dont register.", in.Id)
	}
	return ret, nil
}
