package model

// CouponCreateUpdateBase 优惠券创建或更新时模型
type CouponCreateUpdateBase struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	GoodsIds   string `json:"goods_ids"`
	CategoryId int    `json:"category_id"`
}

// CouponCreateInput 优惠券创建
type CouponCreateInput struct {
	CouponCreateUpdateBase
}

// CouponCreateOutput 优惠券创建
type CouponCreateOutput struct {
	CouponId int64 `json:"coupon_id"`
}

// CouponUpdateInput 优惠券更新
type CouponUpdateInput struct {
	Id int
	CouponCreateUpdateBase
}

// CouponListInput 优惠券列表
type CouponListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// CouponListItem 优惠券列表
type CouponListItem struct {
	Id int `json:"id"`
	CouponCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CouponListItemOutput struct {
	List  []*CouponListItem `json:"list"`
	Total int               `json:"total"`
}
