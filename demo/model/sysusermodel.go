package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sysUserFieldNames          = builderx.FieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "user_id", "create_time", "update_time"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "user_id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SysUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUser struct {
		UserId    int64     `db:"user_id"`
		NickName  string    `db:"nick_name"`
		Phone     string    `db:"phone"`
		RoleId    int64     `db:"role_id"`
		Salt      string    `db:"salt"`
		Avatar    string    `db:"avatar"`
		Sex       string    `db:"sex"`
		Email     string    `db:"email"`
		DeptId    int64     `db:"dept_id"`
		PostId    int64     `db:"post_id"`
		CreateBy  string    `db:"create_by"`
		UpdateBy  string    `db:"update_by"`
		Remark    string    `db:"remark"`
		Status    string    `db:"status"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
	}
)

func NewSysUserModel(conn sqlx.SqlConn, table string) *SysUserModel {
	return &SysUserModel{
		conn:  conn,
		table: table,
	}
}

func (m *SysUserModel) Insert(data SysUser) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + sysUserRowsExpectAutoSet + `) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return m.conn.Exec(query, data.NickName, data.Phone, data.RoleId, data.Salt, data.Avatar, data.Sex, data.Email, data.DeptId, data.PostId, data.CreateBy, data.UpdateBy, data.Remark, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Username, data.Password)
}

func (m *SysUserModel) FindOne(userId int64) (*SysUser, error) {
	query := `select ` + sysUserRows + ` from ` + m.table + ` where user_id = ? limit 1`
	var resp SysUser
	err := m.conn.QueryRow(&resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *SysUserModel) Update(data SysUser) error {
	query := `update ` + m.table + ` set ` + sysUserRowsWithPlaceHolder + ` where user_id = ?`
	_, err := m.conn.Exec(query, data.NickName, data.Phone, data.RoleId, data.Salt, data.Avatar, data.Sex, data.Email, data.DeptId, data.PostId, data.CreateBy, data.UpdateBy, data.Remark, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Username, data.Password, data.UserId)
	return err
}

func (m *SysUserModel) Delete(userId int64) error {
	query := `delete from ` + m.table + ` where user_id = ?`
	_, err := m.conn.Exec(query, userId)
	return err
}
