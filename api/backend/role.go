package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type RoleAddReq struct {
	g.Meta `path:"/backend/role/add" method:"post" tags:"角色管理" summery:"创建角色"`
	Name   string `json:"name" v:"required#角色名不能为空" dc:"角色名"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleAddRes struct {
	RoleId int64 `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"put" tags:"角色管理" summery:"更新角色"`
	Id     int    `json:"id" v:"required|min:1#角色id不能为空|角色id最小为1" dc:"角色id"`
	Name   string `json:"name" v:"required#角色名不能为空" dc:"角色名"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleUpdateRes struct {
	RoleId int64 `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete/{Id}" method:"delete" tags:"角色管理" summery:"删除角色"`
	Id     int `json:"id" v:"required|min:1#角色id不能为空|角色id最小为1" dc:"角色id"`
}

type RoleDeleteRes struct {
	RoleId int64 `json:"role_id"`
}

type RoleListReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"角色管理" summery:"角色列表"`
	common.PaginationReq
}

type RoleListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

type RoleAddPermissionReq struct {
	g.Meta       `path:"/backend/role/add/permission" method:"post" tags:"角色管理" summery:"添加角色权限"`
	RoleId       int `json:"role_id" dc:"角色id"`
	PermissionId int `json:"permission_id" dc:"权限id"`
}

type RoleAddPermissionRes struct {
	Id int64 `json:"id"`
}

type RoleDeletePermissionReq struct {
	g.Meta       `path:"/backend/role/delete/permission" method:"delete" tags:"角色管理" summery:"删除角色权限"`
	RoleId       int `json:"role_id" dc:"角色id"`
	PermissionId int `json:"permission_id" dc:"权限id"`
}

type RoleDeletePermissionRes struct {
}
