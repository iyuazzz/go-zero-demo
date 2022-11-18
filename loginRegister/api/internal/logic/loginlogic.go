package logic

import (
	"context"
	"fmt"
	"loginRegister/model"

	"loginRegister/api/internal/svc"
	"loginRegister/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	user, err := l.svcCtx.UserModel.FindOneByName(context.Background(), req.Name)
	if err != nil {
		if err == model.ErrNotFound {
			err = fmt.Errorf("wrong user name, cant find")
		}
		return nil, err
	}
	if user.Pwd != req.Pwd {
		return nil, fmt.Errorf("wrong password for user (%s)", req.Name)
	}
	resp = &types.LoginRes{
		Ret: fmt.Sprintf("success Login! hello %s ~", user.Name),
	}
	return
}
