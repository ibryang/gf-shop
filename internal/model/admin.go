package model

// AdminCreateUpdateBase 用户创建或更新时模型
type AdminCreateUpdateBase struct {
	Name     string
	Password string
	RoleIds  []int
	UserSalt string
	IsAdmin  int
}

// AdminCreateInput 用户创建时的输入参数
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// AdminCreateOutput 用户创建时的输出参数
type AdminCreateOutput struct {
	RequestId int64 `json:"request_id"`
}

// AdminUpdateInput 用户更新时的输入参数
type AdminUpdateInput struct {
	Id int
	AdminCreateUpdateBase
}

// AdminListInput 用户列表时的输入参数
type AdminListInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// AdminItemOutput 用户列表时的输出参数
type AdminItemOutput struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
	//Password  string `json:"password"`
	RoleIds []int `json:"role_ids"`
	//UserSalt      string `json:"salt"`
	IsAdmin   int    `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminItemListOutput struct {
	List  []*AdminItemOutput `json:"list"`
	Total int                `json:"total"`
}
