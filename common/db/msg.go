package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/groupmessage"
	"imlogic/ent/privatemessage"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/base"
)

type dbMsgModel interface {
	// 添加一个私聊消息
	AddOnePrivateMsg(ctx context.Context, msgType int32, msgId string, chatId, senderId, timestamp int64, content *base.Message) (err error)
	// 添加一个群聊消息
	AddOneGroupMsg(ctx context.Context, msgType int32, msgId string, groupId, senderId, timestamp int64, content *base.Message) (err error)
	// 获取私聊最后一条消息
	GetLastPrivateMsg(ctx context.Context, chatId int64) (msg *ent.PrivateMessage, err error)
	// 获取群聊最后一条消息
	GetLastGroupMsg(ctx context.Context, groupId int64) (msg *ent.GroupMessage, err error)
}

// 获取群聊最后一条消息
func (m *customModel) GetLastGroupMsg(ctx context.Context, groupId int64) (msg *ent.GroupMessage, err error) {
	msg = &ent.GroupMessage{}
	msg, err = m.client.GroupMessage.Query().
		Where(groupmessage.GroupIDEQ(groupId)).
		Order(ent.Desc(groupmessage.FieldTimestamp)).
		First(ctx)
	switch {
	case ent.IsNotFound(err):
		return
	case err != nil:
		return nil, xerr.DbErr(err, "获取群聊最后一条消息失败")
	}
	return
}

// 获取私聊最后一条消息
func (m *customModel) GetLastPrivateMsg(ctx context.Context, chatId int64) (msg *ent.PrivateMessage, err error) {
	msg = &ent.PrivateMessage{}
	msg, err = m.client.PrivateMessage.Query().
		Where(privatemessage.ChatIDEQ(chatId)).
		Order(ent.Desc(privatemessage.FieldTimestamp)).
		First(ctx)
	switch {
	case ent.IsNotFound(err):
		return
	case err != nil:
		return nil, xerr.DbErr(err, "获取私聊最后一条消息失败")
	}
	return

}

// 添加一个群聊消息
func (m *customModel) AddOneGroupMsg(ctx context.Context, msgType int32, msgId string, groupId, senderId, timestamp int64, content *base.Message) (err error) {
	_, err = m.client.GroupMessage.Create().
		SetMsgType(msgType).
		SetMsgID(msgId).
		SetGroupID(groupId).
		SetSenderID(senderId).
		SetTimestamp(timestamp).
		SetContent(content).Save(ctx)
	if err != nil {
		return xerr.DbErr(err, "添加群聊消息失败")
	}
	return
}

// 添加一个私聊消息
func (m *customModel) AddOnePrivateMsg(ctx context.Context, msgType int32, msgId string, chatId, senderId, timestamp int64, content *base.Message) (err error) {
	_, err = m.client.PrivateMessage.Create().
		SetMsgType(msgType).
		SetMsgID(msgId).
		SetChatID(chatId).
		SetSenderID(senderId).
		SetTimestamp(timestamp).
		SetContent(content).Save(ctx)
	if err != nil {
		return xerr.DbErr(err, "添加私聊消息失败")
	}
	return
}
