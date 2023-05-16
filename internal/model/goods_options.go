package model

// GoodsOptionsCreateUpdateBase 商品规格创建或更新时模型
type GoodsOptionsCreateUpdateBase struct {
	GoodsId int    `json:"goods_id" dc:"商品id"`
	Name    string `json:"name" dc:"商品规格名称"`
	Price   int    `json:"price" dc:"价格 单位分"`
	Brand   string `json:"brand" dc:"品牌"`
	PicUrl  string `json:"pic_url" dc:"图片"`
	Stock   int    `json:"stock" dc:"库存"`
}

// GoodsOptionsCreateInput 商品规格创建
type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
}

// GoodsOptionsCreateOutput 商品规格创建
type GoodsOptionsCreateOutput struct {
	Id int64 `json:"id"`
}

// GoodsOptionsUpdateInput 商品规格更新
type GoodsOptionsUpdateInput struct {
	Id int
	GoodsOptionsCreateUpdateBase
}

// GoodsOptionsListInput 商品规格列表
type GoodsOptionsListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GoodsOptionsListItem 商品规格列表
type GoodsOptionsListItem struct {
	Id int `json:"id"`
	GoodsOptionsCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	//entity.GoodsOptionsInfo
}

type GoodsOptionsListItemOutput struct {
	List  []*GoodsOptionsListItem `json:"list"`
	Total int                     `json:"total"`
}
