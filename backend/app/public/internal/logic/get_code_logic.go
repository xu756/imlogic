package logic

import (
	"context"
	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/xu756/imlogic/internal/pb"
	"github.com/xu756/imlogic/internal/xerr"
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

type captcha struct {
	Id string `json:"id"`
}

// GetCode 获取验证码
func (l *GetCodeLogic) GetCode(in *pb.GetCodeReq) (*pb.GetCodeResp, error) {

	id, b64s, err := l.svcCtx.Captcha.Generate()
	if err != nil {
		return nil, xerr.NewSystemError("生成验证码失败" + err.Error())
	}
	return &pb.GetCodeResp{
		Img:       b64s,
		Expire:    120,
		SessionId: id,
	}, nil
}
