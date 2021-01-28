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
	toolTagMapFieldNames          = builderx.RawFieldNames(&ToolTagMap{})
	toolTagMapRows                = strings.Join(toolTagMapFieldNames, ",")
	toolTagMapRowsExpectAutoSet   = strings.Join(stringx.Remove(toolTagMapFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	toolTagMapRowsWithPlaceHolder = strings.Join(stringx.Remove(toolTagMapFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ToolTagMapModel interface {
		Insert(data ToolTagMap) (sql.Result, error)
		FindOne(id int64) (*ToolTagMap, error)
		Update(data ToolTagMap) error
		Delete(id int64) error
	}

	defaultToolTagMapModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ToolTagMap struct {
		Id         int64     `db:"id"`
		ToolId     int64     `db:"tool_id"`     // 工具id
		TagId      int64     `db:"tag_id"`      // 标签id
		CreatedOn  time.Time `db:"created_on"`  // 新建时间
		ModifiedOn time.Time `db:"modified_on"` // 修改时间
		IsDel      string    `db:"is_del"`      // 是否删除 0未删除 1删除
		CreatedBy  string    `db:"created_by"`  // 创建人
		ModifiedBy string    `db:"modified_by"` // 修改人
	}
)

func NewToolTagMapModel(conn sqlx.SqlConn) ToolTagMapModel {
	return &defaultToolTagMapModel{
		conn:  conn,
		table: "`tool_tag_map`",
	}
}

func (m *defaultToolTagMapModel) Insert(data ToolTagMap) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, toolTagMapRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ToolId, data.TagId, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultToolTagMapModel) FindOne(id int64) (*ToolTagMap, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", toolTagMapRows, m.table)
	var resp ToolTagMap
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

func (m *defaultToolTagMapModel) Update(data ToolTagMap) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, toolTagMapRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ToolId, data.TagId, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultToolTagMapModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
