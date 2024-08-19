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

func (s *PublicSrvImpl) GetUserFriendList(ctx context.Context, getOneReq *base.GetOneReq) (r []*user.Friend, err error) {
	r = make([]*user.Friend, 0)
	friends, err := s.Model.GetFriendList(ctx, getOneReq.Id)
	if err != nil {
		return nil, err
	}
	for _, friend := range friends {
		r = append(r, &user.Friend{
			UserId:    friend.WithID,
			Alias:     friend.Alias,
			OwnerDesc: friend.OwnerDesc,
			CreatedAt: friend.CreatedAt.Unix(),
		})
	}
	return r, nil
}
