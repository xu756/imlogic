package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/xu756/imlogic/common/model"
	"github.com/xu756/imlogic/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error)
	FindUserByMobile(ctx context.Context, mobile string) (user *model.User, err error)
	CreateUser(ctx context.Context, userName, mobile string, creator uint64) (err error)
}

// FindUserByUsername 根据用户名查找用户
func (m *customModel) FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error) {
	tx := m.Db.WithContext(ctx)
	err = tx.Model(&model.User{}).Where("user_name = ?", userName).Limit(1).Find(&user).Error
	switch {
	case user.ID == 0:
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.ErrMsg(xerr.DbFindErr)
	}
	return user, nil
}

// FindUserByMobile 根据手机号查找用户
func (m *customModel) FindUserByMobile(ctx context.Context, mobile string) (user *model.User, err error) {
	tx := m.Db.WithContext(ctx)
	err = tx.Model(&model.User{}).Where("mobile = ?", mobile).Limit(1).Find(&user).Error
	switch {
	case user.ID == 0:
		return nil, xerr.ErrMsg(xerr.UserMobileNotExist)
	case err != nil:
		return nil, xerr.ErrMsg(xerr.DbFindErr)
	}
	return user, nil
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, userName, mobile string, creator uint64) (err error) {
	tx := m.Db.WithContext(ctx)
	err = tx.Create(&model.User{
		UserUuid: uuid.NewString(),
		UserName: userName,
		Mobile:   mobile,
		Creator:  creator,
		Editor:   creator,
	}).Error
	if err != nil {
		return xerr.ErrMsg(xerr.UserExist)
	}
	return nil
}
