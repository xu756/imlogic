package db

import (
	"context"
	"fmt"
	"imlogic/ent"
	"imlogic/ent/userfriend"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/base"
)

type dbFriendModel interface {
	// 获取用户聊天列表
	GetUserChatList(ctx context.Context, userId int64) (chatList []*base.ChatList, err error)
	// 判断是不是好友
	CheckIsFriend(ctx context.Context, sender, receiver int64) (err error)
	// 添加好友
	AddOneFriend(ctx context.Context, owner, withId int64) (err error)
	// 获取用户的所有好友
	GetFriendList(ctx context.Context, userId int64) (friends []*ent.UserFriend, err error)
}

// 获取用户的所有好友
func (m *customModel) GetFriendList(ctx context.Context, userId int64) (friends []*ent.UserFriend, err error) {
	friends, err = m.client.UserFriend.Query().
		Where(userfriend.Owner(userId)).
		Order(ent.Desc(userfriend.FieldCreatedAt)).
		All(ctx)
	switch {
	case ent.IsNotFound(err):
		return friends, nil
	case err != nil:
		return nil, xerr.DbErr(err, "获取用户私聊聊天列表失败 用户ID:%d", userId)
	default:
		return friends, nil
	}
}

// 添加好友
func (m *customModel) AddOneFriend(ctx context.Context, owner, withId int64) (err error) {
	// 验证是否是好友
	_, err = m.client.UserFriend.Query().Where(userfriend.Owner(owner), userfriend.WithID(withId)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		userInfo, err := m.GetOneUserById(ctx, withId)
		if err != nil {
			return err
		}
		_, err = m.client.UserFriend.Create().
			SetOwner(owner).
			SetWithID(withId).
			SetAlias(userInfo.Username).
			SetOwnerDesc("").
			SetAgree(true).
			Save(ctx)
		switch {
		case err != nil:
			return xerr.DbErr(err, "添加好友失败")
		}
		_, err = m.client.UserFriend.Create().
			SetOwner(withId).
			SetWithID(owner).
			SetAlias(userInfo.Username).
			SetOwnerDesc("").
			SetAgree(false).
			Save(ctx)
		switch {
		case err != nil:
			return xerr.DbErr(err, "添加好友失败")
		}
		return nil
	case err != nil:
		return xerr.DbErr(err, "判断是不是好友失败")
	default:
		userFriend, err := m.client.UserFriend.Query().Where(userfriend.Owner(withId), userfriend.WithID(owner)).First(ctx)
		if err != nil {
			return xerr.DbErr(err, "判断是不是好友失败")
		}
		if userFriend.Agree {
			return xerr.WarnMsg("已经是好友")
		} else {
			return xerr.WarnMsg("已经发送好友请求")
		}
	}

}

// 判断是不是好友
func (m *customModel) CheckIsFriend(ctx context.Context, sender, receiver int64) (err error) {
	userFriend, err := m.client.UserFriend.Query().Where(userfriend.Owner(receiver), userfriend.WithID(sender)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return xerr.WarnMsg("不是好友")
	case err != nil:
		return xerr.DbErr(err, "判断是不是好友失败")
	default:
		if userFriend.Agree {
			return nil
		}
		return xerr.WarnMsg("对方未同意好友请求")
	}
}

// 获取用户聊天列表
func (m *customModel) GetUserChatList(ctx context.Context, userId int64) (chatList []*base.ChatList, err error) {
	chatList = make([]*base.ChatList, 0)
	// 获取用户私聊聊天列表
	userFriends, err := m.GetFriendList(ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, friendInfo := range userFriends {
		msg, err := m.GetLastPrivateMsg(ctx, userId)
		if err != nil {
			continue
		}
		chatList = append(chatList, &base.ChatList{
			ChatId:    fmt.Sprintf("private-%d", friendInfo.ID),
			ChatType:  base.ChatType_PrivateChat,
			WithId:    friendInfo.WithID,
			LastMsg:   msg.Content,
			Timestamp: msg.Timestamp,
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
		msg, err := m.GetLastGroupMsg(ctx, group.ID)
		if err != nil {
			continue
		}
		chatList = append(chatList, &base.ChatList{
			ChatId:    fmt.Sprintf("group-%d", group.ID),
			ChatType:  base.ChatType_GroupChat,
			WithId:    group.ID,
			LastMsg:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}
	return chatList, nil
}
