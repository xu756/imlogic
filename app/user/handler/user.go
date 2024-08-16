package handler

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/user"
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
func (s *PublicSrvImpl) GetUserOnlineStatus(ctx context.Context, req *base.GetOneReq) (res *user.UserOnlineStatus, err error) {
	status, err := s.Model.GetOneUserStatus(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserOnlineStatus{
		Status: status,
	}, nil
}
