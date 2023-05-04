// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} //
	Username  interface{} // 用户名
	Password  interface{} // 密码
	RoleIds   interface{} // 角色ids
	Sort      interface{} // 排序
	Salt      interface{} // 盐
	IsAdmin   interface{} // 是否是超级管理员
	CreatedAt interface{} //
	UpdatedAt interface{} //
	DeletedAt interface{} //
}