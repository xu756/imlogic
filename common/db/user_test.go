package db

import (
	"context"
	"github.com/xu756/imlogic/common/config"
	"testing"
)

// 根据用户名查找用户
func TestCustomModel_FindUserByUsername(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	userInfo, err := m.FindUserByUsername(context.Background(), "test", "mini")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(userInfo)
}

// 根据手机号查找用户
func TestCustomModel_FindUserByMobile(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	userInfo, err := m.FindUserByMobile(context.Background(), "12345678101", "mini")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(userInfo)
}

// 创建用户
func TestCustomModel_CreateUser(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	user, err := m.CreateUser(context.TODO(), "test", "test", "12345678902", "mini", 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

// 更新用户头像
func TestCustomModel_UpdateUserAvatar(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	m := NewModel()
	user, err := m.UpdateUserAvatar(context.Background(), "4c630aed-387f-4ebb-a3ba-ed40bb44f5a1", "test", "mini", 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}
