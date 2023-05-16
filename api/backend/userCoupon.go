package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type UserCouponAddUpdateBaseReq struct {
	UserId   int `json:"user_id" v:"required#用户id不能为空" dc:"用户id"`
	CouponId int `json:"coupon_id" v:"required#优惠券id不能为空" dc:"优惠券id"`
	Status   int `json:"status" dc:"状态"`
}

type UserCouponReq struct {
	g.Meta `path:"/backend/user/coupon/add " tags:"用户优惠券" method:"post" summary:"添加优惠券"`
	UserCouponAddUpdateBaseReq
}

type UserCouponRes struct {
	Id int64 `json:"id" example:"1" dc:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/backend/user/coupon/delete/{Id}" tags:"用户优惠券" method:"delete" summary:"删除优惠券"`
	Id     int `json:"id" v:"required|min:1#优惠券id不能为空|优惠券id最小为1" dc:"优惠券id"`
}

type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/backend/user/coupon/update/{Id}" tags:"用户优惠券" method:"put" summary:"更新优惠券"`
	Id     int `json:"id" v:"required|min:1#优惠券id不能为空|优惠券id最小为1" dc:"优惠券id"`
	UserCouponAddUpdateBaseReq
}

type UserCouponUpdateRes struct{}

type UserCouponListReq struct {
	g.Meta `path:"/backend/user/coupon/list" tags:"用户优惠券" method:"get" summary:"优惠券列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type UserCouponListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
