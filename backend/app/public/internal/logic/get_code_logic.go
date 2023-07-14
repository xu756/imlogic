package logic

import (
	"context"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCode 获取验证码
func (l *GetCodeLogic) GetCode(in *pb.GetCodeReq) (*pb.GetCodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCodeResp{}, nil
}
