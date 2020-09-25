package logic

import (
	"context"
	
	"EBook/internal/cudr"
	"EBook/internal/model"
	
	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdataUserSecertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdataUserSecertLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdataUserSecertLogic {
	return UpdataUserSecertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdataUserSecertLogic) UpdataUserSecert(req types.UpdataSecertReq) (*types.UpdataSecertResp, error) {
	// 此时已经获取到了jwt但是在更改密码的时候还是需要使用jwt鉴权
	u := &model.User{Username: req.Username, Password: req.Password}
	newp := req.NewPassword
	err := cudr.UpdateSecert(u, newp, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	// 改过密码之后重新生成token
	//jwtTokenResp, err := jwtAuth.JWT(l.svcCtx.Config.JA.AccessExpire, l.svcCtx.Config.JA.AccessSecret)
	//if err != nil {
	//	return nil, err
	//}
	return &types.UpdataSecertResp{Username: u.Username}, nil
}
