package logic

import (
	"context"
	"fmt"
	"strconv"

	"moment/rpc-model/moment-rpc/internal/svc"
	"moment/rpc-model/moment-rpc/types/moment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentLogic {
	return &GetMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMomentLogic) GetMoment(in *moment.GetMomentRequest) (*moment.GetMomentReply, error) {
	if in.Token != "iyua" {
		return nil, fmt.Errorf("wrong token")
	}
	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	ret, err := l.svcCtx.MysqlModel.FindOne(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &moment.GetMomentReply{
		Title:  ret.Title,
		Text:   ret.Text,
		Result: "success",
	}, nil
}
