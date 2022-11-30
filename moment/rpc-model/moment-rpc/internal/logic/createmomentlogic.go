package logic

import (
	"context"
	"fmt"
	mysql_model "moment/mysql-model"
	"moment/rpc-model/moment-rpc/internal/svc"
	"moment/rpc-model/moment-rpc/types/moment"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMomentLogic {
	return &CreateMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMomentLogic) CreateMoment(in *moment.CreateMomentRequest) (*moment.CreateMomentReply, error) {
	if in.Token != "iyua" {
		return nil, fmt.Errorf("wrong token")
	}
	data := &mysql_model.Moment{
		Text:  in.Text,
		Title: in.Title,
	}
	ret, err := l.svcCtx.MysqlModel.Insert(context.Background(), data)
	if err != nil {
		return nil, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &moment.CreateMomentReply{
		Result: strconv.FormatInt(id, 10),
	}, nil
}
