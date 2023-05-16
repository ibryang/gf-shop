package user_coupon

import (
	"context"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// List 用户优惠券列表
func (c *Controller) List(ctx context.Context, req *backend.UserCouponListReq) (res *backend.UserCouponListRes, err error) {
	output, err := service.UserCoupon().List(ctx, &model.UserCouponListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.UserCouponListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 用户优惠券创建
func (c *Controller) Create(ctx context.Context, req *backend.UserCouponReq) (res *backend.UserCouponRes, err error) {
	output, err := service.UserCoupon().Create(ctx, &model.UserCouponCreateInput{
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.UserCouponRes{
		Id: output.Id,
	}
	return
}

// Delete 用户优惠券删除
func (c *Controller) Delete(ctx context.Context, req *backend.UserCouponDeleteReq) (res *backend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, req.Id)
	return
}

// Update 用户优惠券更新
func (c *Controller) Update(ctx context.Context, req *backend.UserCouponUpdateReq) (res *backend.UserCouponUpdateRes, err error) {
	err = service.UserCoupon().Update(ctx, &model.UserCouponUpdateInput{
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
		Id: req.Id,
	})
	return
}
