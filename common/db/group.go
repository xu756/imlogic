package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/group"
	"imlogic/ent/usergroup"
	"imlogic/internal/xerr"
)

type dbGroupModel interface {
	// 根据群ID查找群信息
	GetGroupByGroupId(ctx context.Context, groupId int64) (groupInfo *ent.Group, err error)
	// 根据群ID查找群成员
	GetGroupUserIdsByGroupId(ctx context.Context, groupId int64) (userIds []int64, err error)
	// 获取用户群ID列表
	GetUserGroupIds(ctx context.Context, userId int64) (groupIds []int64, err error)
	// 获取用户群列表
	GetUserGroups(ctx context.Context, groupIds []int64) (groups []*ent.Group, err error)
}

// 获取用户群列表
func (m *customModel) GetUserGroups(ctx context.Context, groupIds []int64) (groups []*ent.Group, err error) {
	groups, err = m.client.Group.Query().
		Where(group.IDIn(groupIds...)).
		All(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "查询用户群列表失败 群ID列表:%v", groupIds)
	}
	return groups, nil
}

// 获取用户群ID列表Ids
func (m *customModel) GetUserGroupIds(ctx context.Context, userId int64) (groupIds []int64, err error) {
	ids, err := m.client.UserGroup.Query().
		Where(usergroup.UserID(userId)).
		IDs(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "查询用户群列表失败 用户ID:%d", userId)
	}

	for _, id := range ids {
		groupIds = append(groupIds, int64(id))
	}
	return groupIds, nil
}

// 根据群ID查找群成员Id
func (m *customModel) GetGroupUserIdsByGroupId(ctx context.Context, groupId int64) (userIds []int64, err error) {
	ids, err := m.client.UserGroup.Query().
		Where(usergroup.GroupID(groupId)).
		IDs(ctx)
	switch {
	case ent.IsNotFound(err):
		return userIds, nil
	case err != nil:
		return nil, xerr.DbErr(err, "查询群成员失败 群ID:%d", groupId)
	}
	userIds = make([]int64, 0, len(ids))
	for _, id := range ids {
		userIds = append(userIds, int64(id))
	}
	return userIds, nil
}

// 根据群ID查找群信息
func (m *customModel) GetGroupByGroupId(ctx context.Context, groupId int64) (groupInfo *ent.Group, err error) {
	groupInfo, err = m.client.Group.Query().
		Where(group.ID(groupId)).
		First(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "查询群信息失败 群ID:%d", groupId)
	}
	return groupInfo, nil
}
