// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"         description:""`
	Name      string      `json:"name"       description:"用户名"`
	Password  string      `json:"password"   description:"密码"`
	RoleIds   string      `json:"role_ids"   description:"角色ids"`
	CreatedAt *gtime.Time `json:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" description:""`
	UserSalt  string      `json:"user_salt"  description:"加密盐"`
	IsAdmin   int         `json:"is_admin"   description:"是否超级管理员"`
}
