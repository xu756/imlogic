package handler

import (
	"context"
	"imlogic/common/cache"
	"imlogic/common/db"
	"imlogic/common/mongodb"
	"imlogic/kitex_gen/im"
	"log"
)

type ImRpcImpl struct {
	Model   db.Model
	Cache   *cache.Client
	mongodb mongodb.Client
}

// AudioMessage implements im.ImSrv.
func (i *ImRpcImpl) AudioMessage(ctx context.Context, req *im.AudioMsg) (res *im.MessageRes, err error) {
	panic("unimplemented")
}

// FileMessage implements im.ImSrv.
func (i *ImRpcImpl) FileMessage(ctx context.Context, req *im.FileMsg) (res *im.MessageRes, err error) {
	panic("unimplemented")
}

// ImageMessage implements im.ImSrv.
func (i *ImRpcImpl) ImageMessage(ctx context.Context, req *im.ImgMsg) (res *im.MessageRes, err error) {
	log.Println("receive image message")
	return &im.MessageRes{}, nil
}

// TextMessage implements im.ImSrv.
func (i *ImRpcImpl) TextMessage(ctx context.Context, req *im.TextMsg) (res *im.MessageRes, err error) {
	log.Println("receive text message")
	log.Print(req)
	return &im.MessageRes{}, nil
}

// VideoMessage implements im.ImSrv.
func (i *ImRpcImpl) VideoMessage(ctx context.Context, req *im.VideoMsg) (res *im.MessageRes, err error) {
	panic("unimplemented")
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model:   db.NewModel(),
		Cache:   cache.NewClient(),
		mongodb: mongodb.NewClient(),
	}
}
