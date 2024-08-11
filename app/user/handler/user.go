package handler

import (
	"context"
	"imlogic/kitex_gen/user"
)

// GetOneUserInfo implements user.UserSrv.
func (s *PublicSrvImpl) GetOneUserInfo(ctx context.Context, req *user.GetOneUserReq) (res *user.UserInfo, err error) {
	userInfo, err := s.Model.GetOneUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserInfo{
		Id:       userInfo.ID,
		Username: userInfo.Username,
		Mobile:   userInfo.Mobile,
		Avatar:   userInfo.Avatar,
		Desc:     userInfo.Desc,
	}, nil
}
