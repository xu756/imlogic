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
}

// GetGroupUserIdsByGroupId 根据群ID查找群成员Id
func (m *customModel) GetGroupUserIdsByGroupId(ctx context.Context, groupId int64) (userIds []int64, err error) {
	ids, err := m.client.UserGroup.Query().
		Where(usergroup.GroupID(groupId)).
		Select(usergroup.FieldUserID).
		IDs(ctx)
	if err != nil {
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
