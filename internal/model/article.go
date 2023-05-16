package model

// ArticleCreateUpdateBase 文章创建或更新时模型
type ArticleCreateUpdateBase struct {
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	PicUrl  string `json:"pic_url"`
	IsAdmin int    `json:"is_admin"`
	Praise  int    `json:"praise"`
	Detail  string `json:"detail"`
}

// ArticleCreateInput 文章创建
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
}

// ArticleCreateOutput 文章创建
type ArticleCreateOutput struct {
	Id int64 `json:"id"`
}

// ArticleUpdateInput 文章更新
type ArticleUpdateInput struct {
	Id int
	ArticleCreateUpdateBase
}

// ArticleListInput 文章列表
type ArticleListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// ArticleListItem 文章列表
type ArticleListItem struct {
	Id int `json:"id"`
	ArticleCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	//entity.ArticleInfo
}

type ArticleListItemOutput struct {
	List  []*ArticleListItem `json:"list"`
	Total int                `json:"total"`
}
