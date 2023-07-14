// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chat

import (
	"context"
	pb2 "github.com/xu756/imlogic/internal/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ImMeta   = pb2.ImMeta
	Message  = pb2.Message
	MetaResp = pb2.MetaResp
	MsgResp  = pb2.MsgResp
	Notice   = pb2.Notice
	Sender   = pb2.Sender

	Chat interface {
		// 元事件 连接 断开 状态更新 解密错误
		Meta(ctx context.Context, in *ImMeta, opts ...grpc.CallOption) (*MetaResp, error)
		// im消息
		Msg(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MsgResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

// 元事件 连接 断开 状态更新 解密错误
func (m *defaultChat) Meta(ctx context.Context, in *ImMeta, opts ...grpc.CallOption) (*MetaResp, error) {
	client := pb2.NewChatClient(m.cli.Conn())
	return client.Meta(ctx, in, opts...)
}

// im消息
func (m *defaultChat) Msg(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MsgResp, error) {
	client := pb2.NewChatClient(m.cli.Conn())
	return client.Msg(ctx, in, opts...)
}