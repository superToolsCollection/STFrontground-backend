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
	tagsFieldNames          = builderx.RawFieldNames(&Tags{})
	tagsRows                = strings.Join(tagsFieldNames, ",")
	tagsRowsExpectAutoSet   = strings.Join(stringx.Remove(tagsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	tagsRowsWithPlaceHolder = strings.Join(stringx.Remove(tagsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	TagsModel interface {
		Insert(data Tags) (sql.Result, error)
		FindOne(id int64) (*Tags, error)
		Update(data Tags) error
		Delete(id int64) error
	}

	defaultTagsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Tags struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`        // 标签名称
		CreatedOn  time.Time `db:"created_on"`  // 新建时间
		ModifiedOn time.Time `db:"modified_on"` // 修改时间
		IsDel      string    `db:"is_del"`      // 是否删除 0未删除 1删除
		State      string    `db:"state"`       // 标签状态 0未激活 1激活
		CreatedBy  string    `db:"created_by"`  // 创建人
		ModifiedBy string    `db:"modified_by"` // 修改人
	}
)

func NewTagsModel(conn sqlx.SqlConn) TagsModel {
	return &defaultTagsModel{
		conn:  conn,
		table: "`tags`",
	}
}

func (m *defaultTagsModel) Insert(data Tags) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tagsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultTagsModel) FindOne(id int64) (*Tags, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tagsRows, m.table)
	var resp Tags
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

func (m *defaultTagsModel) Update(data Tags) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tagsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultTagsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
