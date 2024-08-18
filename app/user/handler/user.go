package handler

import (
	"context"
	"imlogic/kitex_gen/base"
)

// GetOneUserInfo implements user.UserSrv.
func (s *PublicSrvImpl) GetOneUserInfo(ctx context.Context, req *base.GetOneReq) (res *base.UserInfo, err error) {
	userInfo, err := s.Model.GetOneUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &base.UserInfo{
		Id:       userInfo.ID,
		Username: userInfo.Username,
		Mobile:   userInfo.Mobile,
		Avatar:   userInfo.Avatar,
		Desc:     userInfo.Desc,
	}, nil
}

// GetUserOnlineStatus implements user.UserSrv.
func (s *PublicSrvImpl) GetUserOnlineStatus(ctx context.Context, getOneReq *base.GetOneReq) (res *base.BoolRes, err error) {
	status, err := s.Model.GetOneUserStatus(ctx, getOneReq.Id)
	if err != nil {
		return nil, err
	}
	return &base.BoolRes{
		Ok: status,
	}, nil
}
