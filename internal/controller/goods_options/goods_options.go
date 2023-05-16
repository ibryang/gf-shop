package goods_options

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

// List 商品规格列表
func (c *Controller) List(ctx context.Context, req *backend.GoodsOptionsListReq) (res *backend.GoodsOptionsListRes, err error) {
	output, err := service.GoodsOptions().List(ctx, &model.GoodsOptionsListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.GoodsOptionsListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 商品规格创建
func (c *Controller) Create(ctx context.Context, req *backend.GoodsOptionsAddReq) (res *backend.GoodsOptionsRes, err error) {
	var data = model.GoodsOptionsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	output, err := service.GoodsOptions().Create(ctx, &data)
	if err != nil {
		return nil, err
	}
	res = &backend.GoodsOptionsRes{
		Id: output.Id,
	}
	return
}

// Delete 商品规格删除
func (c *Controller) Delete(ctx context.Context, req *backend.GoodsOptionsDeleteReq) (res *backend.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().Delete(ctx, req.Id)
	return
}

// Update 商品规格更新
func (c *Controller) Update(ctx context.Context, req *backend.GoodsOptionsUpdateReq) (res *backend.GoodsOptionsUpdateRes, err error) {
	var data = model.GoodsOptionsUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.GoodsOptions().Update(ctx, &data)
	if err != nil {
		return nil, err
	}
	return
}
