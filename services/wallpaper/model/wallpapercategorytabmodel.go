package model

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	wallpaperCategoryTabFieldNames          = builder.RawFieldNames(&WallpaperCategoryTab{})
	wallpaperCategoryTabRows                = strings.Join(wallpaperCategoryTabFieldNames, ",")
	wallpaperCategoryTabRowsExpectAutoSet   = strings.Join(stringx.Remove(wallpaperCategoryTabFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	wallpaperCategoryTabRowsWithPlaceHolder = strings.Join(stringx.Remove(wallpaperCategoryTabFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWallpaperCategoryTabIdPrefix = "cache:wallpaperCategoryTab:id:"
)

type (
	WallpaperCategoryTabModel interface {
		Insert(data *WallpaperCategoryTab) (sql.Result, error)
		FindOne(id int64) (*WallpaperCategoryTab, error)
		FindList(start, limit int64) ([]*WallpaperCategoryTab, int64, error)
		Update(data *WallpaperCategoryTab) error
		Delete(id int64) error
	}

	defaultWallpaperCategoryTabModel struct {
		sqlc.CachedConn
		table string
	}

	WallpaperCategoryTab struct {
		Id         int64  `db:"id"`          // id
		Category   string `db:"category"`    // category
		Desc       string `db:"desc"`        // desc
		DelFlag    string `db:"del_flag"`    // del flag（0-normal 1-delete)
		CreateTime int64  `db:"create_time"` // create time
		UpdateTime int64  `db:"update_time"` // update time
	}
)

func NewWallpaperCategoryTabModel(conn sqlx.SqlConn, c cache.CacheConf) WallpaperCategoryTabModel {
	return &defaultWallpaperCategoryTabModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`wallpaper_category_tab`",
	}
}

func (m *defaultWallpaperCategoryTabModel) Insert(data *WallpaperCategoryTab) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wallpaperCategoryTabRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Category, data.Desc, data.DelFlag)

	return ret, err
}

func (m *defaultWallpaperCategoryTabModel) FindOne(id int64) (*WallpaperCategoryTab, error) {
	wallpaperCategoryTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperCategoryTabIdPrefix, id)
	var resp WallpaperCategoryTab
	err := m.QueryRow(&resp, wallpaperCategoryTabIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperCategoryTabRows, m.table)
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

func (m *defaultWallpaperCategoryTabModel) FindList(start, limit int64) ([]*WallpaperCategoryTab, int64, error) {
	var resp []*WallpaperCategoryTab
	var query string
	var err error

	if start >= 0 && limit > 0 {
		query = fmt.Sprintf("select %s from %s limit ?, ?", wallpaperCategoryTabRows, m.table)
		err = m.QueryRowsNoCache(&resp, query, start, limit)
	} else {
		query = fmt.Sprintf("select %s from %s", wallpaperCategoryTabRows, m.table)
		err = m.QueryRowsNoCache(&resp, query)
	}

	switch err {
	case nil:
		var total int64
		total, err = m.GetTableCount()
		if err != nil {
			return nil, 0, err
		}
		return resp, total, nil
	case sqlc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		logx.Errorf("FindList, err: %v\n", err)
		return nil, 0, err
	}
}

func (m *defaultWallpaperCategoryTabModel) Update(data *WallpaperCategoryTab) error {
	wallpaperCategoryTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperCategoryTabIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wallpaperCategoryTabRowsWithPlaceHolder)
		return conn.Exec(query, data.Category, data.Desc, data.DelFlag, data.Id)
	}, wallpaperCategoryTabIdKey)
	return err
}

func (m *defaultWallpaperCategoryTabModel) Delete(id int64) error {

	wallpaperCategoryTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperCategoryTabIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, wallpaperCategoryTabIdKey)
	return err
}

func (m *defaultWallpaperCategoryTabModel) GetTableCount() (int64, error) {
	var resp int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		logx.Errorf("GetTableCount, err: %v\n", err)
		return 0, err
	}
}

func (m *defaultWallpaperCategoryTabModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWallpaperCategoryTabIdPrefix, primary)
}

func (m *defaultWallpaperCategoryTabModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperCategoryTabRows, m.table)
	return conn.QueryRow(v, query, primary)
}
