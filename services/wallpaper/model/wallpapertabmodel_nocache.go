package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	NoCacheWallpaperTabModel interface {
		FindList(tid, cid, start, limit int64) ([]*WallpaperTab, int64, error)
		BulkInsert(data []*WallpaperTab) error
		GetTableCount() (int64, error)
		GetTableMaxID() (int64, error)
	}

	noCacheWallpaperTabModel struct {
		sqlx.SqlConn
		table string
	}
)

func NewNoCacheWallpaperTabModel(conn sqlx.SqlConn) NoCacheWallpaperTabModel {
	return &noCacheWallpaperTabModel{
		SqlConn: conn,
		table:   "`wallpaper_tab`",
	}
}

func (m *noCacheWallpaperTabModel) BulkInsert(data []*WallpaperTab) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, wallpaperTabRowsExpectAutoSet)
	bulkInserter, err := sqlx.NewBulkInserter(m.SqlConn, query)
	if err != nil {
		return fmt.Errorf("NewBulkInserter %s", err)
	}
	for k, v := range data {
		if err = bulkInserter.Insert(v.Wid, v.Name, v.Tid, v.Cid, v.ImageUrl, v.Author, v.Desc, v.DelFlag); err != nil {
			return fmt.Errorf("insert k:%d, err:%s", k, err)
		}
	}
	bulkInserter.Flush()
	return nil
}

func (m *noCacheWallpaperTabModel) FindList(tid, cid, start, limit int64) ([]*WallpaperTab, int64, error) {
	var resp []*WallpaperTab
	var query string
	var err error

	if tid == 0 && cid == 0 {
		query = fmt.Sprintf("select %s from %s limit ?, ?", wallpaperTabRows, m.table)
		err = m.QueryRows(&resp, query, start, limit)
	} else if tid == 0 && cid != 0 {
		query = fmt.Sprintf("select %s from %s where cid = ? limit ?, ?", wallpaperTabRows, m.table)
		err = m.QueryRows(&resp, query, cid, start, limit)
	} else if tid != 0 && cid == 0 {
		query = fmt.Sprintf("select %s from %s where tid = ? limit ?, ?", wallpaperTabRows, m.table)
		err = m.QueryRows(&resp, query, tid, start, limit)
	} else {
		query = fmt.Sprintf("select %s from %s where tid = ? and cid = ? limit ?, ?", wallpaperTabRows, m.table)
		err = m.QueryRows(&resp, query, tid, cid, start, limit)
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

func (m *noCacheWallpaperTabModel) GetTableCount() (int64, error) {
	var resp int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.QueryRow(&resp, query)
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

func (m *noCacheWallpaperTabModel) GetTableMaxID() (int64, error) {
	var resp int64
	query := fmt.Sprintf("select coalesce(max(id), 0) from %s", m.table)
	err := m.QueryRow(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		logx.Errorf("GetTableMaxID, err: %v\n", err)
		return 0, err
	}
}
