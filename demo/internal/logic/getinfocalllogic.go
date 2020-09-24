package logic

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetInfoCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInfoCallLogic {
	return GetInfoCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoCallLogic) GetInfoCall(req types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	if req.Username == "lichao" && req.Password == "123456" {
		resp.Code = 200
		resp.Userdata = "你好" + req.Username
		resp.UserInfo = "你是" + req.Username + "哈哈哈"
	}
	return
}
