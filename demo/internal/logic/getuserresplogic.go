package logic

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserRespLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRespLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserRespLogic {
	return GetUserRespLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRespLogic) GetUserResp(req types.UserResp) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	resp.UserInfo = "hhhhhhsdsd"
	resp.Userdata = "ffvdvdfv"
	resp.Code = 200
	return
}
