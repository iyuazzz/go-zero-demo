package logic

import (
	"context"
	"fmt"

	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_tokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet_tokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_tokenLogic {
	return &Get_tokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_tokenLogic) Get_token(req *types.GetTokenReq) (resp *types.GetTokenRep, err error) {
	if req.UserId != "iyua" {
		return nil, fmt.Errorf("fail to get token, user(%s) dont register", req.UserId)
	}
	return &types.GetTokenRep{
		Token:  "iyua",
		Result: "success",
	}, nil
}
