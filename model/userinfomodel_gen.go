// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserInfoIdPrefix    = "cache:userInfo:id:"
	cacheUserInfoEmailPrefix = "cache:userInfo:email:"
	cacheUserInfoNamePrefix  = "cache:userInfo:name:"
)

type (
	userInfoModel interface {
		Insert(ctx context.Context, data *UserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserInfo, error)
		FindOneByEmail(ctx context.Context, email string) (*UserInfo, error)
		FindOneByName(ctx context.Context, name string) (*UserInfo, error)
		Update(ctx context.Context, data *UserInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	UserInfo struct {
		Id         int64     `db:"id"`          // 用户id,自增,唯一主键
		Name       string    `db:"name"`        // 用户昵称
		Password   string    `db:"password"`    // 用户密码
		Email      string    `db:"email"`       // 用户邮箱
		Gender     int64     `db:"gender"`      // 用户性别,0表示未知,1表示男,2表示女
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
		IsDeleted  int64     `db:"is_deleted"`  // 逻辑删除,默认为0,表示未删除,1表示删除
		AvatarUrl  string    `db:"avatar_url"`  // 用户头像
	}
)

func newUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserInfoModel {
	return &defaultUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_info`",
	}
}

func (m *defaultUserInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userInfoEmailKey := fmt.Sprintf("%s%v", cacheUserInfoEmailPrefix, data.Email)
	userInfoIdKey := fmt.Sprintf("%s%v", cacheUserInfoIdPrefix, id)
	userInfoNameKey := fmt.Sprintf("%s%v", cacheUserInfoNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userInfoEmailKey, userInfoIdKey, userInfoNameKey)
	return err
}

func (m *defaultUserInfoModel) FindOne(ctx context.Context, id int64) (*UserInfo, error) {
	userInfoIdKey := fmt.Sprintf("%s%v", cacheUserInfoIdPrefix, id)
	var resp UserInfo
	err := m.QueryRowCtx(ctx, &resp, userInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
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

func (m *defaultUserInfoModel) FindOneByEmail(ctx context.Context, email string) (*UserInfo, error) {
	userInfoEmailKey := fmt.Sprintf("%s%v", cacheUserInfoEmailPrefix, email)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, userInfoEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindOneByName(ctx context.Context, name string) (*UserInfo, error) {
	userInfoNameKey := fmt.Sprintf("%s%v", cacheUserInfoNamePrefix, name)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, userInfoNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) Insert(ctx context.Context, data *UserInfo) (sql.Result, error) {
	userInfoEmailKey := fmt.Sprintf("%s%v", cacheUserInfoEmailPrefix, data.Email)
	userInfoIdKey := fmt.Sprintf("%s%v", cacheUserInfoIdPrefix, data.Id)
	userInfoNameKey := fmt.Sprintf("%s%v", cacheUserInfoNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Password, data.Email, data.Gender, data.IsDeleted, data.AvatarUrl)
	}, userInfoEmailKey, userInfoIdKey, userInfoNameKey)
	return ret, err
}

func (m *defaultUserInfoModel) Update(ctx context.Context, newData *UserInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userInfoEmailKey := fmt.Sprintf("%s%v", cacheUserInfoEmailPrefix, data.Email)
	userInfoIdKey := fmt.Sprintf("%s%v", cacheUserInfoIdPrefix, data.Id)
	userInfoNameKey := fmt.Sprintf("%s%v", cacheUserInfoNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Email, newData.Gender, newData.IsDeleted, newData.AvatarUrl, newData.Id)
	}, userInfoEmailKey, userInfoIdKey, userInfoNameKey)
	return err
}

func (m *defaultUserInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserInfoIdPrefix, primary)
}

func (m *defaultUserInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserInfoModel) tableName() string {
	return m.table
}