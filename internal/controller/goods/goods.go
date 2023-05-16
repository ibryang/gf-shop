package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// List 商品列表
func (c *Controller) List(ctx context.Context, req *backend.GoodsListReq) (res *backend.GoodsListRes, err error) {
	output, err := service.Goods().List(ctx, &model.GoodsListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.GoodsListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 商品创建
func (c *Controller) Create(ctx context.Context, req *backend.GoodsAddReq) (res *backend.GoodsRes, err error) {
	var data = model.GoodsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	output, err := service.Goods().Create(ctx, &data)
	if err != nil {
		return nil, err
	}
	res = &backend.GoodsRes{
		Id: output.Id,
	}
	return
}

// Delete 商品删除
func (c *Controller) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

// Update 商品更新
func (c *Controller) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	var data = model.GoodsUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, &data)
	if err != nil {
		return nil, err
	}
	return
}
