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
	usersAuthFieldNames          = builderx.RawFieldNames(&UsersAuth{})
	usersAuthRows                = strings.Join(usersAuthFieldNames, ",")
	usersAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(usersAuthFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(usersAuthFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UsersAuthModel interface {
		Insert(data UsersAuth) (sql.Result, error)
		FindOne(id int64) (*UsersAuth, error)
		FindOneByPhone(phone string) (*UsersAuth, error)
		FindOneByEmail(email string) (*UsersAuth, error)
		FindOneByName(name string) (*UsersAuth, error)
		Update(data UsersAuth) error
		Delete(id int64) error
	}

	defaultUsersAuthModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UsersAuth struct {
		Id            int64     `db:"id"`
		UserId        int64     `db:"user_id"`        // 用户id
		Name          string    `db:"name"`           // 用户名
		Password      string    `db:"password"`       // 用户密码
		Token         string    `db:"token"`          // 用户token
		Phone         string    `db:"phone"`          // 手机号
		Email         string    `db:"email"`          // 邮箱
		EmailActivate string    `db:"email_activate"` // 邮箱是否激活 0未激活 1激活
		WechatId      string    `db:"wechat_id"`      // 微信id
		WechatToken   string    `db:"wechat_token"`   // 微信token
		WechatExpires string    `db:"wechat_expires"` // 微信token
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
)

func NewUsersAuthModel(conn sqlx.SqlConn) UsersAuthModel {
	return &defaultUsersAuthModel{
		conn:  conn,
		table: "`users_auth`",
	}
}

func (m *defaultUsersAuthModel) Insert(data UsersAuth) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersAuthRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserId, data.Name, data.Password, data.Token, data.Phone, data.Email, data.EmailActivate, data.WechatId, data.WechatToken, data.WechatExpires)
	return ret, err
}

func (m *defaultUsersAuthModel) FindOne(id int64) (*UsersAuth, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersAuthRows, m.table)
	var resp UsersAuth
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

func (m *defaultUsersAuthModel) FindOneByPhone(phone string) (*UsersAuth, error) {
	var resp UsersAuth
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", usersAuthRows, m.table)
	err := m.conn.QueryRow(&resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersAuthModel) FindOneByName(name string) (*UsersAuth, error) {
	var resp UsersAuth
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", usersAuthRows, m.table)
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersAuthModel) FindOneByEmail(email string) (*UsersAuth, error) {
	var resp UsersAuth
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", usersAuthRows, m.table)
	err := m.conn.QueryRow(&resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersAuthModel) Update(data UsersAuth) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersAuthRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserId, data.Name, data.Password, data.Token, data.Phone, data.Email, data.EmailActivate, data.WechatId, data.WechatToken, data.WechatExpires, data.Id)
	return err
}

func (m *defaultUsersAuthModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
