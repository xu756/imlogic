package db

import (
	"context"
	"errors"
	"github.com/xu756/imlogic/common/model"
	"github.com/xu756/imlogic/internal/xerr"
	"gorm.io/gorm"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error)
}

func (m *customModel) FindUserByUsername(ctx context.Context, userName string) (user *model.User, err error) {
	tx := m.Db.WithContext(ctx)
	err = tx.Model(&model.User{}).Where("userName = ?", userName).First(&user).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.ErrMsg(xerr.DbFindErr)
	}
	return user, nil
}
