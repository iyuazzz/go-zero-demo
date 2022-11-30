package logic

import (
	"context"
	"moment/rpc-model/moment-rpc/momentclient"

	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Del_momentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDel_momentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Del_momentLogic {
	return &Del_momentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Del_momentLogic) Del_moment(req *types.DelMomentReq) (resp *types.DelMomentRep, err error) {
	ret, err := l.svcCtx.MomentRpc.DelMoment(context.Background(), &momentclient.DelMomentRequest{
		Id:    req.Id,
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &types.DelMomentRep{
		Result: ret.Result,
	}, nil
}
