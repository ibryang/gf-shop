package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type PermissionAddUpdateBase struct {
	Name string `json:"name" v:"required#权限名不能为空" dc:"权限名"`
	Path string `json:"path" v:"required#路径不能为空" dc:"权限描述"`
}

type PermissionAddReq struct {
	g.Meta `path:"/backend/permission/add" method:"post" tags:"权限管理" summery:"创建权限"`
	PermissionAddUpdateBase
}

type PermissionAddRes struct {
	PermissionId int64 `json:"permission_id"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update" method:"put" tags:"权限管理" summery:"更新权限"`
	Id     int `json:"id" v:"required|min:1#权限id不能为空|权限id最小为1" dc:"权限id"`
	PermissionAddUpdateBase
}

type PermissionUpdateRes struct {
	PermissionId int64 `json:"permission_id"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete/{Id}" method:"delete" tags:"权限管理" summery:"删除权限"`
	Id     int `json:"id" v:"required|min:1#权限id不能为空|权限id最小为1" dc:"权限id"`
}

type PermissionDeleteRes struct {
	PermissionId int64 `json:"permission_id"`
}

type PermissionListReq struct {
	g.Meta `path:"/backend/permission/list" method:"get" tags:"权限管理" summery:"权限列表"`
	common.PaginationReq
}

type PermissionListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
