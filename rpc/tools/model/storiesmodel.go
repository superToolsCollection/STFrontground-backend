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
	storiesFieldNames          = builderx.RawFieldNames(&Stories{})
	storiesRows                = strings.Join(storiesFieldNames, ",")
	storiesRowsExpectAutoSet   = strings.Join(stringx.Remove(storiesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	storiesRowsWithPlaceHolder = strings.Join(stringx.Remove(storiesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	StoriesModel interface {
		Insert(data Stories) (sql.Result, error)
		FindOne(id int64) (*Stories, error)
		Update(data Stories) error
		Delete(id int64) error
	}

	defaultStoriesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Stories struct {
		Id         int64     `db:"id"`
		Auth       string    `db:"auth"`        // 作者
		Story      string    `db:"story"`       // 故事
		CreatedOn  time.Time `db:"created_on"`  // 新建时间
		ModifiedOn time.Time `db:"modified_on"` // 修改时间
		IsDel      string    `db:"is_del"`      // 是否删除 0未删除 1删除
		State      string    `db:"state"`       // 故事状态 0未激活 1激活
		CreatedBy  string    `db:"created_by"`  // 创建人
		ModifiedBy string    `db:"modified_by"` // 修改人
	}
)

func NewStoriesModel(conn sqlx.SqlConn) StoriesModel {
	return &defaultStoriesModel{
		conn:  conn,
		table: "`stories`",
	}
}

func (m *defaultStoriesModel) Insert(data Stories) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, storiesRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Auth, data.Story, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultStoriesModel) FindOne(id int64) (*Stories, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storiesRows, m.table)
	var resp Stories
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

func (m *defaultStoriesModel) Update(data Stories) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, storiesRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Auth, data.Story, data.CreatedOn, data.ModifiedOn, data.IsDel, data.State, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultStoriesModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
