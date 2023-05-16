package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type CouponAddUpdateBaseReq struct {
	Name       string `json:"name" v:"required#优惠券名称不能为空" dc:"优惠券名称"`
	Price      int    `json:"price" v:"required#优惠券金额不能为空" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" v:"" dc:"可用的商品ids, 逗号分割"`
	CategoryId int    `json:"coupon_id" dc:"可用的商品分类"`
}

type CouponReq struct {
	g.Meta `path:"/backend/coupon/add" tags:"优惠券" method:"post" summary:"添加优惠券"`
	CouponAddUpdateBaseReq
}

type CouponRes struct {
	CouponId int64 `json:"coupon_id" example:"1" dc:"coupon id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/backend/coupon/delete/{Id}" tags:"优惠券" method:"delete" summary:"删除优惠券"`
	Id     int `json:"id" v:"required|min:1#优惠券id不能为空|优惠券id最小为1" dc:"优惠券id"`
}

type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/backend/coupon/update/{Id}" tags:"优惠券" method:"put" summary:"更新优惠券"`
	Id     int `json:"id" v:"required|min:1#优惠券id不能为空|优惠券id最小为1" dc:"优惠券id"`
	CouponAddUpdateBaseReq
}

type CouponUpdateRes struct{}

type CouponListReq struct {
	g.Meta `path:"/backend/coupon/list" tags:"优惠券" method:"get" summary:"优惠券列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type CouponListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
