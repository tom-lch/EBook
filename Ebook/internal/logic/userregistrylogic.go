package logic

import (
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

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

	return &types.UserRegistryResp{}, nil
}
