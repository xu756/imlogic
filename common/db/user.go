package db

import (
	"context"
	"imlogic/common/types"
	"imlogic/ent"
	"imlogic/ent/chat"
	"imlogic/ent/user"
	"imlogic/internal/xerr"
)

type dbUserModel interface {
	// 通过用户名查找用户
	GetOneUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error)
	// 通过手机号查找用户
	GetOneUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error)
	// 通过 Ids 查找用户
	GetUsersByIds(ctx context.Context, ids []int64) (users []*ent.User, err error)
	// 通过id查找用户
	GetOneUserById(ctx context.Context, id int64) (userInfo *ent.User, err error)
	// 获取用户是否在线
	GetOneUserStatus(ctx context.Context, id int64) (status bool, err error)
}

// ChatList 聊天列表
type ChatList struct {
	Uuid     string         `json:"uuid"`
	ChatType types.ChatType `json:"chat_type"`
	GroupId  int64          `json:"group_id"`
	// 用户1 id
	User1Id int64 `json:"user1_id"`
	// 用户2 id
	User2Id int64 `json:"user2_id"`
	// 最后一条消息
	LastMsgId string `json:"last_msg_id"`
	// 最后一条消息时间
	Timestamp int64 `json:"timestamp"`
}

// 获取用户聊天列表
func (m *customModel) GetUserChatList(ctx context.Context, userId int64) (chatList []*ChatList, err error) {
	chatList = make([]*ChatList, 0)
	// 获取用户私聊聊天列表
	chatListEnts, err := m.client.Chat.Query().
		Where(
			chat.Or(
				chat.User1ID(userId),
				chat.User2ID(userId),
			),
		).
		Order(ent.Desc(chat.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "获取用户私聊聊天列表失败 用户ID:%d", userId)
	}
	for _, chatEnt := range chatListEnts {
		chatList = append(chatList, &ChatList{
			Uuid:     chatEnt.UUID,
			ChatType: types.PrivateChat,
			User1Id:  chatEnt.User1ID,
			User2Id:  chatEnt.User2ID,
		})
	}
	// 获取用户群聊聊天列表
	groupIds, err := m.GetUserGroupIds(ctx, userId)
	if err != nil {
		return nil, xerr.DbErr(err, "获取用户群聊聊天列表失败 用户ID:%d", userId)
	}
	groups, err := m.GetUserGroups(ctx, groupIds)
	if err != nil {
		return nil, xerr.DbErr(err, "获取用户群聊聊天列表失败 用户ID:%d", userId)
	}
	for _, group := range groups {
		chatList = append(chatList, &ChatList{
			Uuid:     group.UUID,
			ChatType: types.GroupChat,
			GroupId:  group.ID,
			User1Id:  userId,
		})
	}

}

// GetOneUserStatus 获取用户是否在线
func (m *customModel) GetOneUserStatus(ctx context.Context, id int64) (status bool, err error) {
	userConns, err := m.GetUserConnsByUserId(ctx, id)
	if err != nil {
		return false, err
	}
	return len(userConns) > 0, nil
}

// GetOneUserById 通过id查找用户
func (m *customModel) GetOneUserById(ctx context.Context, id int64) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.ID(id), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 用户ID:%d", id)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 用户ID:%d", id)
	default:
		return userInfo, nil
	}
}

// GetUsersByIds 通过 Ids 查找用户
func (m *customModel) GetUsersByIds(ctx context.Context, ids []int64) (users []*ent.User, err error) {
	users, err = m.client.User.Query().
		Where(user.IDIn(ids...), user.Deleted(false)).All(ctx)
	if err != nil {
		return nil, xerr.DbErr(err, "查询用户失败 用户ID:%v", ids)
	}
	return users, nil
}

// 根据用户名查找用户
func (m *customModel) GetOneUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Username(username), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 用户名:%s", username)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 用户名:%s", username)
	default:
		return userInfo, nil
	}
}

// 根据手机号查找用户
func (m *customModel) GetOneUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Mobile(mobile), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 手机号:%s", mobile)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 手机号:%s", mobile)
	default:
		return userInfo, nil
	}
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile string, editor int64) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Create().
		SetUsername(username).
		SetPassword(password).
		SetMobile(mobile).
		SetEditor(editor).
		Save(ctx)
	switch {
	case ent.IsConstraintError(err):
		return nil, xerr.WarnMsg("用户已存在 用户名:%s 手机号:%s", username, mobile)
	case err != nil:
		return nil, xerr.DbErr(err, "创建用户失败 用户名:%s 手机号:%s", username, mobile)
	default:
		return userInfo, nil
	}
}
