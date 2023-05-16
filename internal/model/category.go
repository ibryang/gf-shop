package model

// CategoryCreateUpdateBase 商品分类创建或更新时模型
type CategoryCreateUpdateBase struct {
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	PicUrl   string `json:"pic_url"`
	Sort     int    `json:"sort"`
	Level    int    `json:"level"`
}

// CategoryCreateInput 商品分类创建时的输入参数
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
}

// CategoryCreateOutput 商品分类创建时的输出参数
type CategoryCreateOutput struct {
	CategoryId int64 `json:"category_id"`
}

// CategoryUpdateInput 商品分类更新时的输入参数
type CategoryUpdateInput struct {
	Id int
	CategoryCreateUpdateBase
}

// CategoryListInput 商品分类列表时的输入参数
type CategoryListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// CategoryListItem 商品分类列表时的输出参数
type CategoryListItem struct {
	Id int `json:"id"`
	CategoryCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryListItemOutput struct {
	List  []*CategoryListItem `json:"list"`
	Total int                 `json:"total"`
}
