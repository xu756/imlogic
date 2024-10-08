// Code generated by Kitex v0.8.0. DO NOT EDIT.

package usersrv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "imlogic/kitex_gen/base"
	user "imlogic/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LoginByPassword(ctx context.Context, loginByPasswordReq *user.LoginByPasswordReq, callOptions ...callopt.Option) (r *user.LoginRes, err error)
	LoginByMobile(ctx context.Context, loginByMobileReq *user.LoginByMobileReq, callOptions ...callopt.Option) (r *user.LoginRes, err error)
	SendCaptcha(ctx context.Context, sendCaptchaReq *user.SendCaptchaReq, callOptions ...callopt.Option) (r *base.BoolRes, err error)
	GetOneUserInfo(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r *base.UserInfo, err error)
	GetUserOnlineStatus(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r *base.BoolRes, err error)
	GetUserChatList(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r []*base.ChatList, err error)
	GetUserFriendList(ctx context.Context, getFriendListReq *user.GetFriendListReq, callOptions ...callopt.Option) (r []*user.Friend, err error)
	AddFriend(ctx context.Context, addFriendReq *user.AddFriendReq, callOptions ...callopt.Option) (r *base.BoolRes, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserSrvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserSrvClient struct {
	*kClient
}

func (p *kUserSrvClient) LoginByPassword(ctx context.Context, loginByPasswordReq *user.LoginByPasswordReq, callOptions ...callopt.Option) (r *user.LoginRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginByPassword(ctx, loginByPasswordReq)
}

func (p *kUserSrvClient) LoginByMobile(ctx context.Context, loginByMobileReq *user.LoginByMobileReq, callOptions ...callopt.Option) (r *user.LoginRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginByMobile(ctx, loginByMobileReq)
}

func (p *kUserSrvClient) SendCaptcha(ctx context.Context, sendCaptchaReq *user.SendCaptchaReq, callOptions ...callopt.Option) (r *base.BoolRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendCaptcha(ctx, sendCaptchaReq)
}

func (p *kUserSrvClient) GetOneUserInfo(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r *base.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOneUserInfo(ctx, getOneReq)
}

func (p *kUserSrvClient) GetUserOnlineStatus(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r *base.BoolRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserOnlineStatus(ctx, getOneReq)
}

func (p *kUserSrvClient) GetUserChatList(ctx context.Context, getOneReq *base.GetOneReq, callOptions ...callopt.Option) (r []*base.ChatList, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserChatList(ctx, getOneReq)
}

func (p *kUserSrvClient) GetUserFriendList(ctx context.Context, getFriendListReq *user.GetFriendListReq, callOptions ...callopt.Option) (r []*user.Friend, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserFriendList(ctx, getFriendListReq)
}

func (p *kUserSrvClient) AddFriend(ctx context.Context, addFriendReq *user.AddFriendReq, callOptions ...callopt.Option) (r *base.BoolRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFriend(ctx, addFriendReq)
}
