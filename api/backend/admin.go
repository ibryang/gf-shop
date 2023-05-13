package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type AdminAddReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"用户" method:"get" summary:"添加用户"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  []int  `json:"role_ids" dc:"角色"`
	IsAdmin  int    `json:"is_admin" description:"is_admin" dc:"是否管理员"`
}

type AdminAddRes struct {
	AdminId int64 `json:"admin_id" example:"1" description:"admin id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete/{Id}" tags:"用户" method:"delete" summary:"删除用户"`
	Id     int `json:"id" v:"required|min:1#用户id不能为空|用户id最小为1" dc:"用户id"`
}

type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" tags:"用户" method:"put" summary:"更新用户"`
	Id       int    `json:"id" v:"required|min:1#用户id不能为空|用户id最小为1" dc:"用户id"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  []int  `json:"role_ids" dc:"角色"`
	IsAdmin  int    `json:"is_admin" description:"is_admin" dc:"是否管理员"`
}

type AdminUpdateRes struct{}

type AdminListReq struct {
	g.Meta `path:"/backend/admin/list" tags:"用户" method:"get" summary:"用户列表"`
	common.PaginationReq
}

type AdminListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

type AdminInfoReq struct {
	g.Meta `path:"/backend/admin/info" tags:"用户" method:"get" summary:"用户信息"`
}

// for jwt
//type AdminInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// for gtoken
type AdminInfoRes struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	RoleIds  []int  `json:"role_ids"`
	IsAdmin  int    `json:"is_admin"`
}
