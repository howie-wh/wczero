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
	wallpaperTabFieldNames          = builder.RawFieldNames(&WallpaperTab{})
	wallpaperTabRows                = strings.Join(wallpaperTabFieldNames, ",")
	wallpaperTabRowsExpectAutoSet   = strings.Join(stringx.Remove(wallpaperTabFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	wallpaperTabRowsWithPlaceHolder = strings.Join(stringx.Remove(wallpaperTabFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWallpaperTabIdPrefix  = "cache:wallpaperTab:id:"
	cacheWallpaperTabWidPrefix = "cache:wallpaperTab:wid:"
	cacheWallpaperTabListPrefix = "cache:wallpaperTab:list:"

	cacheExpire = time.Minute * 5
)

type (
	WallpaperTabModel interface {
		Insert(data *WallpaperTab) (sql.Result, error)
		FindOne(id int64) (*WallpaperTab, error)
		FindOneByWid(wid string) (*WallpaperTab, error)
		FindList(start, limit int64) ([]*WallpaperTab, error)
		Update(data *WallpaperTab) error
		Delete(id int64) error
		DeleteByWid(wid string) error
	}
	NoCacheWallpaperTabModel interface {
		BulkInsert(data []*WallpaperTab) error
	}

	defaultWallpaperTabModel struct {
		sqlc.CachedConn
		table string
	}

	noCacheWallpaperTabModel struct {
		sqlx.SqlConn
		table string
	}

	WallpaperTab struct {
		Id         int64  `db:"id"`          // id
		Wid        string `db:"wid"`         // wallpaper id
		Name       string `db:"name"`        // name
		ImageUrl   string `db:"image_url"`   // image url
		Author     string `db:"author"`      // author
		Desc       string `db:"desc"`        // desc
		DelFlag    string `db:"del_flag"`    // del flag（0-normal 1-delete)
		CreateTime int64  `db:"create_time"` // create time
		UpdateTime int64  `db:"update_time"` // update time
	}
)

func NewWallpaperTabModel(conn sqlx.SqlConn, c cache.CacheConf) WallpaperTabModel {
	return &defaultWallpaperTabModel{
		CachedConn: sqlc.NewConn(conn, c, cache.WithExpiry(cacheExpire)),
		table:      "`wallpaper_tab`",
	}
}

func NewNoCacheWallpaperTabModel(conn sqlx.SqlConn) NoCacheWallpaperTabModel {
	return &noCacheWallpaperTabModel{
		SqlConn: conn,
		table:   "`wallpaper_tab`",
	}
}

func (m *defaultWallpaperTabModel) Insert(data *WallpaperTab) (sql.Result, error) {
	wallpaperTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, data.Id)
	wallpaperTabWidKey := fmt.Sprintf("%s%v", cacheWallpaperTabWidPrefix, data.Wid)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, wallpaperTabRowsExpectAutoSet)
		return conn.Exec(query, data.Wid, data.Name, data.ImageUrl, data.Author, data.Desc, data.DelFlag)
	}, wallpaperTabIdKey, wallpaperTabWidKey)
	return ret, err
}

func (m *noCacheWallpaperTabModel) BulkInsert(data []*WallpaperTab) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, wallpaperTabRowsExpectAutoSet)
	bulkInserter, err := sqlx.NewBulkInserter(m.SqlConn, query)
	if err != nil{
		return fmt.Errorf("NewBulkInserter %s",err)
	}
	for k,v := range data{
		if err = bulkInserter.Insert(v.Wid, v.Name, v.ImageUrl, v.Author, v.Desc, v.DelFlag);err != nil{
			fmt.Println("insert",err,k)
			return fmt.Errorf("insert k:%d, err:%s", k, err)
		}
	}
	bulkInserter.Flush()
	return nil
}

func (m *defaultWallpaperTabModel) FindOne(id int64) (*WallpaperTab, error) {
	wallpaperTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, id)
	var resp WallpaperTab
	err := m.QueryRow(&resp, wallpaperTabIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperTabRows, m.table)
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

func (m *defaultWallpaperTabModel) FindOneByWid(wid string) (*WallpaperTab, error) {
	wallpaperTabWidKey := fmt.Sprintf("%s%v", cacheWallpaperTabWidPrefix, wid)
	var resp WallpaperTab
	err := m.QueryRowIndex(&resp, wallpaperTabWidKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `wid` = ? limit 1", wallpaperTabRows, m.table)
		if err := conn.QueryRow(&resp, query, wid); err != nil {
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

func (m *defaultWallpaperTabModel) FindList(start, limit int64) ([]*WallpaperTab, error) {
	wallpaperTabListKey := fmt.Sprintf("%s%d:%d", cacheWallpaperTabListPrefix, start, limit)
	resp := make([]*WallpaperTab, 0)
	err := m.QueryRow(resp, wallpaperTabListKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where limit ?, ?", wallpaperTabRows, m.table)
		return conn.QueryRow(resp, query, start, limit)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWallpaperTabModel) Update(data *WallpaperTab) error {
	wallpaperTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, data.Id)
	wallpaperTabWidKey := fmt.Sprintf("%s%v", cacheWallpaperTabWidPrefix, data.Wid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wallpaperTabRowsWithPlaceHolder)
		return conn.Exec(query, data.Wid, data.Name, data.ImageUrl, data.Author, data.Desc, data.DelFlag, data.Id)
	}, wallpaperTabIdKey, wallpaperTabWidKey)
	return err
}

func (m *defaultWallpaperTabModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	wallpaperTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, id)
	wallpaperTabWidKey := fmt.Sprintf("%s%v", cacheWallpaperTabWidPrefix, data.Wid)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, wallpaperTabIdKey, wallpaperTabWidKey)
	return err
}

func (m *defaultWallpaperTabModel) DeleteByWid(wid string) error {
	data, err := m.FindOneByWid(wid)
	if err != nil {
		return err
	}

	wallpaperTabIdKey := fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, data.Id)
	wallpaperTabWidKey := fmt.Sprintf("%s%v", cacheWallpaperTabWidPrefix, data.Wid)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `wid` = ?", m.table)
		return conn.Exec(query, wid)
	}, wallpaperTabIdKey, wallpaperTabWidKey)
	return err
}

func (m *defaultWallpaperTabModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWallpaperTabIdPrefix, primary)
}

func (m *defaultWallpaperTabModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wallpaperTabRows, m.table)
	return conn.QueryRow(v, query, primary)
}
