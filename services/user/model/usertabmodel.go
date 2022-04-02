package model

import (
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
	userTabFieldNames          = builder.RawFieldNames(&UserTab{})
	userTabRows                = strings.Join(userTabFieldNames, ",")
	userTabRowsExpectAutoSet   = strings.Join(stringx.Remove(userTabFieldNames, "`user_id`", "`create_time`", "`update_time`"), ",")
	userTabRowsWithPlaceHolder = strings.Join(stringx.Remove(userTabFieldNames, "`user_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserTabUserIdPrefix   = "cache:userTab:userId:"
	cacheUserTabUserNamePrefix = "cache:userTab:userName:"
)

type (
	UserTabModel interface {
		Insert(data *UserTab) (sql.Result, error)
		FindOne(userId int64) (*UserTab, error)
		FindOneByUserName(userName string) (*UserTab, error)
		Update(data *UserTab) error
		Delete(userId int64) error
	}

	defaultUserTabModel struct {
		sqlc.CachedConn
		table string
	}

	UserTab struct {
		UserId     int64  `db:"user_id"`     // id
		UserName   string `db:"user_name"`   // username
		NickName   string `db:"nick_name"`   // nickname
		Email      string `db:"email"`       // email
		Avatar     string `db:"avatar"`      // avatar
		DelFlag    string `db:"del_flag"`    // del flag（0-normal 1-delete)
		CreateTime int64  `db:"create_time"` // create time
		UpdateTime int64  `db:"update_time"` // update time
	}
)

func NewUserTabModel(conn sqlx.SqlConn, c cache.CacheConf) UserTabModel {
	return &defaultUserTabModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_tab`",
	}
}

func (m *defaultUserTabModel) Insert(data *UserTab) (sql.Result, error) {
	userTabUserIdKey := fmt.Sprintf("%s%v", cacheUserTabUserIdPrefix, data.UserId)
	userTabUserNameKey := fmt.Sprintf("%s%v", cacheUserTabUserNamePrefix, data.UserName)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userTabRowsExpectAutoSet)
		return conn.Exec(query, data.UserName, data.NickName, data.Email, data.Avatar, data.DelFlag)
	}, userTabUserIdKey, userTabUserNameKey)
	return ret, err
}

func (m *defaultUserTabModel) FindOne(userId int64) (*UserTab, error) {
	userTabUserIdKey := fmt.Sprintf("%s%v", cacheUserTabUserIdPrefix, userId)
	var resp UserTab
	err := m.QueryRow(&resp, userTabUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userTabRows, m.table)
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

func (m *defaultUserTabModel) FindOneByUserName(userName string) (*UserTab, error) {
	userTabUserNameKey := fmt.Sprintf("%s%v", cacheUserTabUserNamePrefix, userName)
	var resp UserTab
	err := m.QueryRowIndex(&resp, userTabUserNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", userTabRows, m.table)
		if err := conn.QueryRow(&resp, query, userName); err != nil {
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

func (m *defaultUserTabModel) Update(data *UserTab) error {
	userTabUserIdKey := fmt.Sprintf("%s%v", cacheUserTabUserIdPrefix, data.UserId)
	userTabUserNameKey := fmt.Sprintf("%s%v", cacheUserTabUserNamePrefix, data.UserName)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userTabRowsWithPlaceHolder)
		return conn.Exec(query, data.UserName, data.NickName, data.Email, data.Avatar, data.DelFlag, data.UserId)
	}, userTabUserIdKey, userTabUserNameKey)
	return err
}

func (m *defaultUserTabModel) Delete(userId int64) error {
	data, err := m.FindOne(userId)
	if err != nil {
		return err
	}

	userTabUserIdKey := fmt.Sprintf("%s%v", cacheUserTabUserIdPrefix, userId)
	userTabUserNameKey := fmt.Sprintf("%s%v", cacheUserTabUserNamePrefix, data.UserName)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.Exec(query, userId)
	}, userTabUserIdKey, userTabUserNameKey)
	return err
}

func (m *defaultUserTabModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserTabUserIdPrefix, primary)
}

func (m *defaultUserTabModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userTabRows, m.table)
	return conn.QueryRow(v, query, primary)
}
