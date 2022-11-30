package logic

import (
	"context"
	"moment/rpc-model/moment-rpc/momentclient"

	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_momentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet_momentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_momentLogic {
	return &Get_momentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_momentLogic) Get_moment(req *types.GetMomentReq) (resp *types.GetMomentRep, err error) {
	rep, err := l.svcCtx.MomentRpc.GetMoment(context.Background(), &momentclient.GetMomentRequest{
		Id:    req.Id,
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetMomentRep{
		Title:  rep.Title,
		Text:   rep.Text,
		Result: rep.Result,
	}, nil
}
