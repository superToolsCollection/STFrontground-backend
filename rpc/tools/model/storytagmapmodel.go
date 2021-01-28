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
	storyTagMapFieldNames          = builderx.RawFieldNames(&StoryTagMap{})
	storyTagMapRows                = strings.Join(storyTagMapFieldNames, ",")
	storyTagMapRowsExpectAutoSet   = strings.Join(stringx.Remove(storyTagMapFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	storyTagMapRowsWithPlaceHolder = strings.Join(stringx.Remove(storyTagMapFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	StoryTagMapModel interface {
		Insert(data StoryTagMap) (sql.Result, error)
		FindOne(id int64) (*StoryTagMap, error)
		Update(data StoryTagMap) error
		Delete(id int64) error
	}

	defaultStoryTagMapModel struct {
		conn  sqlx.SqlConn
		table string
	}

	StoryTagMap struct {
		Id         int64     `db:"id"`
		StoryId    int64     `db:"story_id"`    // 睡前故事ID
		TagId      int64     `db:"tag_id"`      // 标签ID
		CreatedOn  time.Time `db:"created_on"`  // 新建时间
		ModifiedOn time.Time `db:"modified_on"` // 修改时间
		IsDel      string    `db:"is_del"`      // 是否删除 0未删除 1删除
		CreatedBy  string    `db:"created_by"`  // 创建人
		ModifiedBy string    `db:"modified_by"` // 修改人
	}
)

func NewStoryTagMapModel(conn sqlx.SqlConn) StoryTagMapModel {
	return &defaultStoryTagMapModel{
		conn:  conn,
		table: "`story_tag_map`",
	}
}

func (m *defaultStoryTagMapModel) Insert(data StoryTagMap) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, storyTagMapRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.StoryId, data.TagId, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy)
	return ret, err
}

func (m *defaultStoryTagMapModel) FindOne(id int64) (*StoryTagMap, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storyTagMapRows, m.table)
	var resp StoryTagMap
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

func (m *defaultStoryTagMapModel) Update(data StoryTagMap) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, storyTagMapRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.StoryId, data.TagId, data.CreatedOn, data.ModifiedOn, data.IsDel, data.CreatedBy, data.ModifiedBy, data.Id)
	return err
}

func (m *defaultStoryTagMapModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
