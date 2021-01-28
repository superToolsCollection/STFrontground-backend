package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	forbesFieldNames          = builderx.RawFieldNames(&Forbes{})
	forbesRows                = strings.Join(forbesFieldNames, ",")
	forbesRowsExpectAutoSet   = strings.Join(stringx.Remove(forbesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	forbesRowsWithPlaceHolder = strings.Join(stringx.Remove(forbesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ForbesModel interface {
		Insert(data Forbes) (sql.Result, error)
		FindOne(id int64) (*Forbes, error)
		Update(data Forbes) error
		Delete(id int64) error
	}

	defaultForbesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Forbes struct {
		Id             int64     `db:"id"`
		Rank           int64     `db:"rank"`             // 排名
		Name           string    `db:"name"`             // 中文名
		NameEn         string    `db:"name_en"`          // 英文名
		Wealth         float64   `db:"wealth"`           // 财富
		SourceOfWealth string    `db:"source_of_wealth"` // 财富来源
		Region         string    `db:"region"`           // 地区
		CreatedOn      time.Time `db:"created_on"`       // 新建时间
		ModifiedOn     time.Time `db:"modified_on"`      // 修改时间
		IsDel          string    `db:"is_del"`           // 是否删除 0未删除 1删除
		CreatedBy      string    `db:"created_by"`       // 创建人
		ModifiedBy     string    `db:"modified_by"`      // 修改人
	}
)

func NewForbesModel(conn sqlx.SqlConn) ForbesModel {
	return &defaultForbesModel{
		conn:  conn,
		table: "`forbes`",
	}
}

func (m *defaultForbesModel) Insert(data Forbes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, forbesRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Rank, data.Name, data.NameEn, data.Wealth, data.SourceOfWealth, data.Region, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultForbesModel) FindOne(id int64) (*Forbes, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", forbesRows, m.table)
	var resp Forbes
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultForbesModel) Update(data Forbes) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, forbesRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Rank, data.Name, data.NameEn, data.Wealth, data.SourceOfWealth, data.Region, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultForbesModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
