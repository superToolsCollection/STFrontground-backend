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
	toolsFieldNames          = builderx.RawFieldNames(&Tools{})
	toolsRows                = strings.Join(toolsFieldNames, ",")
	toolsRowsExpectAutoSet   = strings.Join(stringx.Remove(toolsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	toolsRowsWithPlaceHolder = strings.Join(stringx.Remove(toolsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ToolsModel interface {
		Insert(data Tools) (sql.Result, error)
		FindOne(id int64) (*Tools, error)
		FindOneByApi(api string) (*Tools, error)
		Update(data Tools) error
		Delete(id int64) error
	}

	defaultToolsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Tools struct {
		Id          int64     `db:"id"`
		Name        string    `db:"name"`         // 工具名称
		Api         string    `db:"api"`          // 工具api链接
		ApiDescribe string    `db:"api_describe"` // 工具描述
		Picture     string    `db:"picture"`      // 工具图标
		CreatedOn   time.Time `db:"created_on"`   // 新建时间
		ModifiedOn  time.Time `db:"modified_on"`  // 修改时间
		IsDel       int64     `db:"is_del"`       // 是否删除 0为未删除 1为已删除
		State       int64     `db:"state"`        // 状态 0为未上线 1为上线
		CreatedBy   string    `db:"created_by"`   // 创建人
		ModifiedBy  string    `db:"modified_by"`  // 修改人
	}
)

func NewToolsModel(conn sqlx.SqlConn) ToolsModel {
	return &defaultToolsModel{
		conn:  conn,
		table: "`tools`",
	}
}

func (m *defaultToolsModel) Insert(data Tools) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, toolsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Api, data.ApiDescribe, data.Picture, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultToolsModel) FindOne(id int64) (*Tools, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", toolsRows, m.table)
	var resp Tools
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

func (m *defaultToolsModel) FindOneByApi(api string) (*Tools, error) {
	var resp Tools
	query := fmt.Sprintf("select %s from %s where `api` = ? limit 1", toolsRows, m.table)
	err := m.conn.QueryRow(&resp, query, api)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultToolsModel) Update(data Tools) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, toolsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Api, data.ApiDescribe, data.Picture, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultToolsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
