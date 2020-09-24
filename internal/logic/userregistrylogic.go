package logic

import (
	"EBook/internal/cudr"
	"EBook/internal/model"
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"
	"EBook/pkg/utils"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserRegistryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegistryLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserRegistryLogic {
	return UserRegistryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegistryLogic) UserRegistry(req types.UserRegistryReq) (*types.UserRegistryResp, error) {
	// todo: add your logic here and delete this line
	// 用户登录 ，首先go-zero帮助我们解析好了http的post表单，以及按照预先声明的方式放入了UserRegistryReq中
	// 我们首先要做的就是验证传输过来的字段是否有效
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
		"Email":    {utils.NotEmpty()},
	}
	if UserVerifyErr := utils.Verify(req, UserVerify); UserVerifyErr != nil {
		logx.Error("该用户注册信息有误, err:", UserVerifyErr)
		return nil, UserVerifyErr
	}
	// 验证完毕无误后，连接数据库进行插入操作
	user := &model.User{Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	err := cudr.Register(*user, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	resp := &types.UserRegistryResp{
		NickName: user.NickName,
		Email:    user.Email,
		Username: user.Username,
	}
	return resp, nil
}
