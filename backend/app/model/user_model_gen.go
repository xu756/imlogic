// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserIdPrefix = "cache:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64  `db:"id"`           // 用户ID
		State      int64  `db:"state"`        // 用户状态:0=正常,1=禁用
		Name       string `db:"name"`         // 姓名
		HeadImgUrl string `db:"head_img_url"` // 头像图片地址
		Mobile     string `db:"mobile"`       // 手机号码
		Salt       string `db:"salt"`         // 密码加盐
		Password   string `db:"password"`     // 登录密码
		Created    int64  `db:"created"`      // 创建时间
		Creator    int64  `db:"creator"`      // 创建人
		Edited     int64  `db:"edited"`       // 修改时间
		Editor     int64  `db:"editor"`       // 修改人
		Deleted    int64  `db:"deleted"`      // 逻辑删除:0=未删除,1=已删除
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) withSession(session sqlx.Session) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.State, data.Name, data.HeadImgUrl, data.Mobile, data.Salt, data.Password, data.Created, data.Creator, data.Edited, data.Editor, data.Deleted)
	}, userIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.State, data.Name, data.HeadImgUrl, data.Mobile, data.Salt, data.Password, data.Created, data.Creator, data.Edited, data.Editor, data.Deleted, data.Id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
