package logic

import (
	"context"
	"moment/rpc-model/moment-rpc/momentclient"

	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Create_momentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate_momentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Create_momentLogic {
	return &Create_momentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create_momentLogic) Create_moment(req *types.CreateMomentReq) (resp *types.CreateMomentRep, err error) {
	rep, err := l.svcCtx.MomentRpc.CreateMoment(context.Background(), &momentclient.CreateMomentRequest{
		Title: req.Title,
		Text:  req.Text,
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateMomentRep{
		Result: rep.Result,
	}, nil
}
