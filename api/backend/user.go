package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type UserReq struct {
	g.Meta   `path:"/backend/user/add " tags:"用户" method:"post" summary:"添加用户"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  []int  `json:"role_ids" dc:"角色"`
	IsAdmin  int    `json:"is_admin" description:"is_admin" dc:"是否管理员"`
	Sort     int    `json:"sort" description:"sort" dc:"排序" d:"1"`
}

type UserRes struct {
	UserId int64 `json:"user_id" example:"1" description:"user id"`
}

type UserDeleteReq struct {
	g.Meta `path:"/backend/user/delete/{Id}" tags:"用户" method:"delete" summary:"删除用户"`
	Id     int `json:"id" v:"required|min:1#用户id不能为空|用户id最小为1" dc:"用户id"`
}

type UserDeleteRes struct{}

type UserUpdateReq struct {
	g.Meta   `path:"/backend/user/update/{Id}" tags:"用户" method:"put" summary:"更新用户"`
	Id       int    `json:"id" v:"required|min:1#用户id不能为空|用户id最小为1" dc:"用户id"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  []int  `json:"role_ids" dc:"角色"`
	IsAdmin  int    `json:"is_admin" description:"is_admin" dc:"是否管理员"`
	Sort     int    `json:"sort" description:"sort" dc:"排序" d:"1"`
}

type UserUpdateRes struct{}

type UserListReq struct {
	g.Meta `path:"/backend/user/list" tags:"用户" method:"get" summary:"用户列表"`
	Sort   int `json:"sort" description:"sort" dc:"排序"`
	common.PaginationReq
}

type UserListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

type UserInfoReq struct {
	g.Meta `path:"/backend/user/info" tags:"用户" method:"get" summary:"用户信息"`
}

type UserInfoRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}
