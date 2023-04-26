package model

// PositionCreateUpdateBase 轮播图创建或更新时模型
type PositionCreateUpdateBase struct {
	PicUrl    string
	Link      string
	Sort      int
	GoodsId   int
	GoodsName string
	Remark    string
	Status    int
}

// PositionCreateInput 轮播图创建时的输入参数
type PositionCreateInput struct {
	PositionCreateUpdateBase
}

// PositionCreateOutput 轮播图创建时的输出参数
type PositionCreateOutput struct {
	RequestId int64 `json:"request_id"`
}

// PositionUpdateInput 轮播图更新时的输入参数
type PositionUpdateInput struct {
	Id int
	PositionCreateUpdateBase
}

// PositionListInput 轮播图列表时的输入参数
type PositionListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// PositionListItem 轮播图列表时的输出参数
type PositionListItem struct {
	Id        int    `json:"id"`
	PicUrl    string `json:"pic_url"`
	Link      string `json:"link"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	GoodsId   int    `json:"goods_id"`
	GoodsName string `json:"goods_name"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PositionListItemOutput struct {
	List  []*PositionListItem `json:"list"`
	Total int                 `json:"total"`
}
