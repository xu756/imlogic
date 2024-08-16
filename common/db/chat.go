package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/chat"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/im"
)

type dbChatModel interface {
	GetUserChatList(ctx context.Context, userId int64) (chatList []*ChatList, err error)
}

// ChatList 聊天列表
type ChatList struct {
	Uuid     string      `json:"uuid"`
	ChatType im.ChatType `json:"chat_type"`
	GroupId  int64       `json:"group_id"`
	ChatId   int64       `json:"chat_id"`
	// 用户1 id
	User1Id int64 `json:"user1_id"`
	// 用户2 id
	User2Id int64 `json:"user2_id"`
	// 最后一条消息
	LastMsg *im.Message `json:"last_msg"`
	// 最后一条消息时间
	Timestamp int64 `json:"timestamp"`
}

// 获取用户聊天列表
func (m *customModel) GetUserChatList(ctx context.Context, userId int64) (chatList []*ChatList, err error) {
	chatList = make([]*ChatList, 0)
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
		chatList = append(chatList, &ChatList{
			Uuid:      chatInfo.UUID,
			ChatType:  im.ChatType_PrivateChat,
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
		chatList = append(chatList, &ChatList{
			Uuid:      group.UUID,
			ChatType:  im.ChatType_GroupChat,
			GroupId:   group.ID,
			User1Id:   userId,
			LastMsg:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}
	return chatList, nil

}
