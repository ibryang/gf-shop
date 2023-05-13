package model

type RoleCreateUpdateBase struct {
	Name string `json:"name" v:"required#角色名不能为空" dc:"角色名"`
	Desc string `json:"desc" dc:"角色描述"`
}

type RoleAddInput struct {
	RoleCreateUpdateBase
}

type RoleAddOutput struct {
	RoleId int64 `json:"role_id"`
}

// RoleUpdateInput 用户更新时的输入参数
type RoleUpdateInput struct {
	Id int
	RoleCreateUpdateBase
}

// RoleListInput 用户列表时的输入参数
type RoleListInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// RoleItemOutput 用户列表时的输出参数
type RoleItemOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RoleItemListOutput struct {
	List  []*RoleItemOutput `json:"list"`
	Total int               `json:"total"`
}

type RoleAddPermissionInput struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

type RoleAddPermissionOutput struct {
	Id int64 `json:"id"`
}

type RoleDeletePermissionInput struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

type RoleDeletePermissionOutput struct {
}
