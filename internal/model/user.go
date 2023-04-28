package model

// UserCreateUpdateBase 用户创建或更新时模型
type UserCreateUpdateBase struct {
	Username string
	Password string
	RoleIds  []int
	Salt     string
	IsAdmin  int
	Sort     int
}

// UserCreateInput 用户创建时的输入参数
type UserCreateInput struct {
	UserCreateUpdateBase
}

// UserCreateOutput 用户创建时的输出参数
type UserCreateOutput struct {
	RequestId int64 `json:"request_id"`
}

// UserUpdateInput 用户更新时的输入参数
type UserUpdateInput struct {
	Id int
	UserCreateUpdateBase
}

// UserListInput 用户列表时的输入参数
type UserListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// UserListItem 用户列表时的输出参数
type UserListItem struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	//Password  string `json:"password"`
	RoleIds []int `json:"role_ids"`
	//Salt      string `json:"salt"`
	IsAdmin   int    `json:"is_admin"`
	Sort      int    `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserListItemOutput struct {
	List  []*UserListItem `json:"list"`
	Total int             `json:"total"`
}
