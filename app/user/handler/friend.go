package handler

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) AddFriend(ctx context.Context, addFriendReq *user.AddFriendReq) (r *base.BoolRes, err error) {
	err = s.Model.AddOneFriend(ctx, addFriendReq.Owner, addFriendReq.WithId)
	if err != nil {
		return nil, err
	}
	return &base.BoolRes{Ok: true}, nil
}
