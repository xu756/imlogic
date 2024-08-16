package handler

import (
	"context"
	"imlogic/kitex_gen/base"
)

func (s *PublicSrvImpl) GetUserChatList(ctx context.Context, getOneReq *base.GetOneReq) (res []*base.ChatList, err error) {
	res, err = s.Model.GetUserChatList(ctx, getOneReq.Id)
	if err != nil {
		return nil, err
	}
	return
}
