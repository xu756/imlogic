package handler

import (
	"context"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) AddFriend(ctx context.Context, addFriendReq *user.AddFriendReq) (r *base.BoolRes, err error) {
	if addFriendReq.Owner == addFriendReq.WithId {
		return &base.BoolRes{Ok: false}, xerr.WarnMsg("不能添加自己为好友")
	}
	err = s.Model.AddOneFriend(ctx, addFriendReq.Owner, addFriendReq.WithId)
	if err != nil {
		return nil, err
	}
	return &base.BoolRes{Ok: true}, nil
}

func (s *PublicSrvImpl) GetUserFriendList(ctx context.Context, getFriendListReq *user.GetFriendListReq) (r []*user.Friend, err error) {
	r = make([]*user.Friend, 0)
	friends, err := s.Model.GetFriendList(ctx, getFriendListReq.Id, getFriendListReq.Agree)
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
