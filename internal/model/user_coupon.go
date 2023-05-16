package model

// UserCouponCreateUpdateBase 用户优惠券创建或更新时模型
type UserCouponCreateUpdateBase struct {
	UserId   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
	Status   int `json:"status"`
}

// UserCouponCreateInput 用户优惠券创建
type UserCouponCreateInput struct {
	UserCouponCreateUpdateBase
}

// UserCouponCreateOutput 用户优惠券创建
type UserCouponCreateOutput struct {
	Id int64 `json:"id"`
}

// UserCouponUpdateInput 用户优惠券更新
type UserCouponUpdateInput struct {
	Id int
	UserCouponCreateUpdateBase
}

// UserCouponListInput 用户优惠券列表
type UserCouponListInput struct {
	Sort     int `json:"sort"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// UserCouponListItem 用户优惠券列表
type UserCouponListItem struct {
	Id int `json:"id"`
	UserCouponCreateUpdateBase
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserCouponListItemOutput struct {
	List  []*UserCouponListItem `json:"list"`
	Total int                   `json:"total"`
}
