package logic

import (
	"context"
	"fmt"
	"strconv"

	"moment/rpc-model/moment-rpc/internal/svc"
	"moment/rpc-model/moment-rpc/types/moment"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMomentLogic {
	return &DelMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMomentLogic) DelMoment(in *moment.DelMomentRequest) (*moment.DelMomentReply, error) {
	if in.Token != "iyua" {
		return nil, fmt.Errorf("wrong token")
	}
	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.MysqlModel.Delete(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &moment.DelMomentReply{
		Result: "success",
	}, nil
}
