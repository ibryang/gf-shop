package model

type PermissionCreateUpdateBase struct {
	Name string `json:"name" v:"required#权限名不能为空" dc:"权限名"`
	Path string `json:"path" v:"required#路径不能为空" dc:"权限描述"`
}

type PermissionAddInput struct {
	PermissionCreateUpdateBase
}

type PermissionAddOutput struct {
	PermissionId int64 `json:"permission_id"`
}

// PermissionUpdateInput 用户更新时的输入参数
type PermissionUpdateInput struct {
	Id int
	PermissionCreateUpdateBase
}

// PermissionListInput 用户列表时的输入参数
type PermissionListInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// PermissionItemOutput 用户列表时的输出参数
type PermissionItemOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type PermissionItemListOutput struct {
	List  []*PermissionItemOutput `json:"list"`
	Total int                     `json:"total"`
}
