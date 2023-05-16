package coupon

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

// List 优惠券列表
func (c *Controller) List(ctx context.Context, req *backend.CouponListReq) (res *backend.CouponListRes, err error) {
	output, err := service.Coupon().List(ctx, &model.CouponListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.CouponListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 优惠券创建
func (c *Controller) Create(ctx context.Context, req *backend.CouponReq) (res *backend.CouponRes, err error) {
	output, err := service.Coupon().Create(ctx, &model.CouponCreateInput{
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
			Price:      req.Price,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.CouponRes{
		CouponId: output.CouponId,
	}
	return
}

// Delete 优惠券删除
func (c *Controller) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.Id)
	return
}

// Update 优惠券更新
func (c *Controller) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, &model.CouponUpdateInput{
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
			Price:      req.Price,
		},
		Id: req.Id,
	})
	return
}
