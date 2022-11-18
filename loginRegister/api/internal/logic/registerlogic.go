package logic

import (
	"context"
	"fmt"
	"loginRegister/model"

	"loginRegister/api/internal/svc"
	"loginRegister/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	resp = &types.RegisterRes{
		Ret: "success",
	}
	if _, err = l.svcCtx.UserModel.FindOneByName(context.Background(), req.Name); err != model.ErrNotFound {
		if err == nil {
			err = fmt.Errorf("id (%s) already exist, please login", req.Name)
		}
		resp = &types.RegisterRes{
			Ret:    "fail",
			Reason: err.Error(),
		}
		return
	}

	data := &model.User{
		Name: req.Name,
		Pwd:  req.Pwd,
	}
	if err = l.svcCtx.UserModel.Insert(context.Background(), data); err != nil {
		resp = &types.RegisterRes{
			Ret:    "fail",
			Reason: err.Error(),
		}
	}
	return
}
