package position

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

// List 手工位列表
func (c *Controller) List(ctx context.Context, req *backend.PositionListReq) (res *backend.PositionListRes, err error) {
	output, err := service.Position().List(ctx, &model.PositionListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.PositionListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 手工位创建
func (c *Controller) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	output, err := service.Position().Create(ctx, &model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsId:   req.GoodsId,
			GoodsName: req.GoodsName,
			Sort:      req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.PositionRes{
		PositionId: output.RequestId,
	}
	return
}

// Delete 手工位删除
func (c *Controller) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

// Update 手工位更新
func (c *Controller) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, &model.PositionUpdateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
			GoodsName: req.GoodsName,
		},
		Id: req.Id,
	})
	return
}
