package model

// GoodsCreateUpdateBase 用户优惠券创建或更新时模型
type GoodsCreateUpdateBase struct {
	PicUrl           string `json:"pic_url"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	Level1CategoryId int    `json:"level1_category_id"`
	Level2CategoryId int    `json:"level2_category_id"`
	Level3CategoryId int    `json:"level3_category_id"`
	Brand            string `json:"brand"`
	Stock            int    `json:"stock"`
	Sale             int    `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
}

// GoodsCreateInput 用户优惠券创建
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 用户优惠券创建
type GoodsCreateOutput struct {
	Id int64 `json:"id"`
}

// GoodsUpdateInput 用户优惠券更新
type GoodsUpdateInput struct {
	Id int
	GoodsCreateUpdateBase
}

// GoodsListInput 用户优惠券列表
type GoodsListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GoodsListItem 用户优惠券列表
type GoodsListItem struct {
	Id int `json:"id"`
	GoodsCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	//entity.GoodsInfo
}

type GoodsListItemOutput struct {
	List  []*GoodsListItem `json:"list"`
	Total int              `json:"total"`
}
