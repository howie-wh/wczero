package model

import (
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
	userAdminTabFieldNames          = builder.RawFieldNames(&UserAdminTab{})
	userAdminTabRows                = strings.Join(userAdminTabFieldNames, ",")
	userAdminTabRowsExpectAutoSet   = strings.Join(stringx.Remove(userAdminTabFieldNames, "`user_id`", "`create_time`", "`update_time`"), ",")
	userAdminTabRowsWithPlaceHolder = strings.Join(stringx.Remove(userAdminTabFieldNames, "`user_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserAdminTabUserIdPrefix = "cache:userAdminTab:userId:"
	cacheUserAdminTabMobilePrefix = "cache:userAdminTab:mobile:"
)

type (
	UserAdminTabModel interface {
		Insert(data *UserAdminTab) (sql.Result, error)
		FindOne(userId int64) (*UserAdminTab, error)
		FindOneByMobile(mobile string) (*UserAdminTab, error)
		Update(data *UserAdminTab) error
		Delete(userId int64) error
	}

	defaultUserAdminTabModel struct {
		sqlc.CachedConn
		table string
	}

	UserAdminTab struct {
		UserId     int64     `db:"user_id"`
		UserName   string    `db:"user_name"` // 用户姓名
		Gender     int64     `db:"gender"`    // 用户性别
		Mobile     string    `db:"mobile"`    // 用户电话
		Password   string    `db:"password"`  // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewUserAdminTabModel(conn sqlx.SqlConn, c cache.CacheConf) UserAdminTabModel {
	return &defaultUserAdminTabModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_admin_tab`",
	}
}

func (m *defaultUserAdminTabModel) Insert(data *UserAdminTab) (sql.Result, error) {
	userAdminTabUserIdKey := fmt.Sprintf("%s%v", cacheUserAdminTabUserIdPrefix, data.UserId)
	userAdminTabMobileKey := fmt.Sprintf("%s%v", cacheUserAdminTabMobilePrefix, data.Mobile)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userAdminTabRowsExpectAutoSet)
		return conn.Exec(query, data.UserName, data.Gender, data.Mobile, data.Password)
	}, userAdminTabUserIdKey, userAdminTabMobileKey)
	return ret, err
}

func (m *defaultUserAdminTabModel) FindOne(userId int64) (*UserAdminTab, error) {
	userAdminTabUserIdKey := fmt.Sprintf("%s%v", cacheUserAdminTabUserIdPrefix, userId)
	var resp UserAdminTab
	err := m.QueryRow(&resp, userAdminTabUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userAdminTabRows, m.table)
		return conn.QueryRow(v, query, userId)
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

func (m *defaultUserAdminTabModel) FindOneByMobile(mobile string) (*UserAdminTab, error) {
	userAdminTabMobileKey := fmt.Sprintf("%s%v", cacheUserAdminTabMobilePrefix, mobile)
	var resp UserAdminTab
	err := m.QueryRowIndex(&resp, userAdminTabMobileKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userAdminTabRows, m.table)
		if err := conn.QueryRow(&resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.UserId, nil
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

func (m *defaultUserAdminTabModel) Update(data *UserAdminTab) error {
	userAdminTabUserIdKey := fmt.Sprintf("%s%v", cacheUserAdminTabUserIdPrefix, data.UserId)
	userAdminTabMobileKey := fmt.Sprintf("%s%v", cacheUserAdminTabMobilePrefix, data.Mobile)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userAdminTabRowsWithPlaceHolder)
		return conn.Exec(query, data.UserName, data.Gender, data.Mobile, data.Password, data.UserId)
	}, userAdminTabUserIdKey, userAdminTabMobileKey)
	return err
}

func (m *defaultUserAdminTabModel) Delete(userId int64) error {
	data, err := m.FindOne(userId)
	if err != nil {
		return err
	}

	userAdminTabUserIdKey := fmt.Sprintf("%s%v", cacheUserAdminTabUserIdPrefix, userId)
	userAdminTabMobileKey := fmt.Sprintf("%s%v", cacheUserAdminTabMobilePrefix, data.Mobile)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.Exec(query, userId)
	}, userAdminTabUserIdKey, userAdminTabMobileKey)
	return err
}

func (m *defaultUserAdminTabModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAdminTabUserIdPrefix, primary)
}

func (m *defaultUserAdminTabModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userAdminTabRows, m.table)
	return conn.QueryRow(v, query, primary)
}
