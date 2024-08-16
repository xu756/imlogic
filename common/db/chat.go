package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/chat"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/base"
)

type dbChatModel interface {
	GetUserChatList(ctx context.Context, userId int64) (chatList []*base.ChatList, err error)
}

// 获取用户聊天列表
func (m *customModel) GetUserChatList(ctx context.Context, userId int64) (chatList []*base.ChatList, err error) {
	chatList = make([]*base.ChatList, 0)
	// 获取用户私聊聊天列表
	chatsInfo, err := m.client.Chat.Query().
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
	for _, chatInfo := range chatsInfo {
		msg, err := m.GetLastPrivateMsg(ctx, chatInfo.ID)
		if err != nil {
			continue
		}
		chatList = append(chatList, &base.ChatList{
			Uuid:      chatInfo.UUID,
			ChatType:  base.ChatType_PrivateChat,
			ChatId:    chatInfo.ID,
			User1Id:   chatInfo.User1ID,
			User2Id:   chatInfo.User2ID,
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
			Uuid:      group.UUID,
			ChatType:  base.ChatType_GroupChat,
			GroupId:   group.ID,
			User1Id:   userId,
			LastMsg:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}
	return chatList, nil

}