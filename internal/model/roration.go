package model

// RotationCreateUpdateBase 轮播图创建或更新时模型
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
	Remark string
	Status int
}

// RotationCreateInput 轮播图创建时的输入参数
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 轮播图创建时的输出参数
type RotationCreateOutput struct {
	RequestId int64 `json:"request_id"`
}

// RotationUpdateInput 轮播图更新时的输入参数
type RotationUpdateInput struct {
	Id int
	RotationCreateUpdateBase
}

// RotationListInput 轮播图列表时的输入参数
type RotationListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// RotationListItem 轮播图列表时的输出参数
type RotationListItem struct {
	Id        int    `json:"id"`
	PicUrl    string `json:"pic_url"`
	Link      string `json:"link"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RotationListItemOutput struct {
	List  []*RotationListItem `json:"list"`
	Total int                 `json:"total"`
}
