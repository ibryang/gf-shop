// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id           int         `json:"id"            description:""`
	Name         string      `json:"name"          description:"用户名"`
	Avatar       string      `json:"avatar"        description:"头像"`
	Password     string      `json:"password"      description:""`
	UserSalt     string      `json:"user_salt"     description:"加密盐 生成密码用"`
	Sex          int         `json:"sex"           description:"1男 2女"`
	Status       int         `json:"status"        description:"1正常 2拉黑冻结"`
	Sign         string      `json:"sign"          description:"个性签名"`
	SecretAnswer string      `json:"secret_answer" description:"密保问题的答案"`
	CreatedAt    *gtime.Time `json:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updated_at"    description:""`
	DeletedAt    *gtime.Time `json:"deleted_at"    description:""`
}
