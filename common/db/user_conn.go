package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/userconn"
	"imlogic/internal/xerr"
)

type dbUserConnModel interface {
	//  根据用户ID查找用户连接信息
	GetUserConnsByUserId(ctx context.Context, userId int64) (userConns []*ent.UserConn, err error)
	//  添加用户连接信息
	AddUserConn(ctx context.Context, userId int64, hostName, device, linkId string) (err error)
	// 获取所有连接信息
	GetAllUserConns(ctx context.Context) (userConns []*ent.UserConn, err error)
}

// 根据用户ID查找用户连接信息
func (m *customModel) GetUserConnsByUserId(ctx context.Context, userId int64) (userConns []*ent.UserConn, err error) {
	userConns, err = m.client.UserConn.Query().
		Where(userconn.UserID(userId)).
		Order(ent.Desc(userconn.FieldLinkTime)).
		All(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户连接信息不存在 用户ID:%d", userId)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户连接信息失败 用户ID:%d", userId)
	}
	return userConns, nil
}

// 添加用户连接信息
func (m *customModel) AddUserConn(ctx context.Context, userId int64, hostName, device, linkId string) (err error) {
	_, err = m.client.UserConn.Create().
		SetUserID(userId).
		SetHostName(hostName).
		SetDevice(device).
		SetLinkID(linkId).
		Save(ctx)
	if err != nil {
		return xerr.DbErr(err, "添加用户连接信息失败 用户ID:%d 主机名:%s 设备:%s 连接ID:%s", userId, hostName, device, linkId)
	}
	return nil
}

// 获取所有连接信息
func (m *customModel) GetAllUserConns(ctx context.Context) (userConns []*ent.UserConn, err error) {
	userConns, err = m.client.UserConn.Query().
		Order(ent.Desc(userconn.FieldLinkTime)).
		All(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "查询所有用户连接信息失败")
	}
	return userConns, nil
}
