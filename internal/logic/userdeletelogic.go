package logic

import (
	"EBook/internal/cudr"
	"EBook/internal/model"
	"EBook/pkg/utils"
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req types.UserDeleteReq) (*types.UserDeleteResp, error) {
	// todo: add your logic here and delete this line
	// 首先验证请求是否合法
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, UserVerify); err != nil {
		logx.Error("验证出错")
		return nil, err
	}
	// 调用数据库删除用户
	u := &model.User{Username: req.Username, Password: req.Password}
	if err := cudr.Delete(u, l.svcCtx.DB); err != nil {
		return nil, err
	}
	return &types.UserDeleteResp{Info: req.Username + ",用户删除道成功"}, nil
}
