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
	wallpaperTypeTabFieldNames          = builder.RawFieldNames(&WallpaperTypeTab{})
	wallpaperTypeTabRows                = strings.Join(wallpaperTypeTabFieldNames, ",")
	wallpaperTypeTabRowsExpectAutoSet   = strings.Join(stringx.Remove(wallpaperTypeTabFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	wallpaperTypeTabRowsWithPlaceHolder = strings.Join(stringx.Remove(wallpaperTypeTabFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWallpaperTypeTabIdPrefix = "cache:wallpaperTypeTab:id:"
)

type (
	WallpaperTypeTabModel interface {
		Insert(data *WallpaperTypeTab) (sql.Result, error)
		FindOne(id int64) (*WallpaperTypeTab, error) //1
		FindList(start, limit int64) ([]*WallpaperTypeTab, int64, error)
		Update(data *WallpaperTypeTab) error
		Delete(id int64) error
	}

	defaultWallpaperTypeTabModel struct {
		sqlc.CachedConn
		table string
	}

	WallpaperTypeTab struct {
		Id         int64  `db:"id"`          // id
		Tp         string `db:"tp"`          // tp
		Desc       string `db:"desc"`        // desc
		DelFlag    string `db:"del_flag"`    // del flag（0-normal 1-delete)
		CreateTime int64  `db:"create_time"` // create time
		UpdateTime int64  `db:"update_time"` // update time
	}
)

func NewWallpaperTypeTabModel(conn sqlx.SqlConn, c cache.CacheConf) WallpaperTypeTabModel {
	return &defaultWallpaperTypeTabModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`wallpaper_type_tab`",
	}
}

func (m *defaultWallpaperTypeTabModel) Insert(data *WallpaperTypeTab) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wallpaperTypeTabRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Tp, data.Desc, data.DelFlag)

	return ret, err
}

func (m *defaultWallpaperTypeTabModel) FindOne(id int64) (*WallpaperTypeTab, error) {
	wallpaperTypeTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTypeTabIdPrefix, id)
	var resp WallpaperTypeTab
	err := m.QueryRow(&resp, wallpaperTypeTabIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperTypeTabRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultWallpaperTypeTabModel) FindList(start, limit int64) ([]*WallpaperTypeTab, int64, error) {
	var resp []*WallpaperTypeTab
	var query string
	var err error

	if start >= 0 && limit > 0 {
		query = fmt.Sprintf("select %s from %s limit ?, ?", wallpaperTypeTabRows, m.table)
		err = m.QueryRowNoCache(&resp, query, start, limit)
	} else {
		query = fmt.Sprintf("select %s from %s", wallpaperTypeTabRows, m.table)
		err = m.QueryRowNoCache(&resp, query)
	}

	switch err {
	case nil:
		//var total int64
		//total, err = m.GetTableCount()
		//if err != nil {
		//	return nil, 0, err
		//}
		return resp, 0, nil
	case sqlc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

func (m *defaultWallpaperTypeTabModel) Update(data *WallpaperTypeTab) error {
	wallpaperTypeTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTypeTabIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wallpaperTypeTabRowsWithPlaceHolder)
		return conn.Exec(query, data.Tp, data.Desc, data.DelFlag, data.Id)
	}, wallpaperTypeTabIdKey)
	return err
}

func (m *defaultWallpaperTypeTabModel) Delete(id int64) error {

	wallpaperTypeTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTypeTabIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, wallpaperTypeTabIdKey)
	return err
}

func (m *defaultWallpaperTypeTabModel) GetTableCount() (int64, error) {
	var resp int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.QueryRowNoCache(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultWallpaperTypeTabModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWallpaperTypeTabIdPrefix, primary)
}

func (m *defaultWallpaperTypeTabModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperTypeTabRows, m.table)
	return conn.QueryRow(v, query, primary)
}
