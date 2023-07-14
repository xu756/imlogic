package logic

import (
	"context"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/api/public/internal/svc"
	"github.com/xu756/imlogic/api/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ByPasswordLogic {
	return &ByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ByPasswordLogic) ByPassword(req *types.LoginReq) (resp *types.LoginRes, err error) {
	result, err := l.svcCtx.PublicRpc.LoginByPassword(l.ctx, &pb.LoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		SessionId: req.SessionId,
		Code:      req.Code,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.LoginRes{
		Expire: result.Expire,
		Token:  result.Token,
	}
	return resp, nil
}
