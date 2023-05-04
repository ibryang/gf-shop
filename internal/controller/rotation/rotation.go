package rotation

import (
	"context"
	"shop/api/backend"
	"shop/api/frontend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// List 轮播图列表
func (c *Controller) List(ctx context.Context, req *backend.RotationListReq) (res *backend.RotationListRes, err error) {
	output, err := service.Rotation().List(ctx, &model.RotationListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.RotationListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// FrontendList 前台轮播图列表
func (c *Controller) FrontendList(ctx context.Context, req *frontend.RotationListReq) (res *frontend.RotationListRes, err error) {
	output, err := service.Rotation().List(ctx, &model.RotationListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &frontend.RotationListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 轮播图创建
func (c *Controller) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	output, err := service.Rotation().Create(ctx, &model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.RotationRes{
		RotationId: output.RequestId,
	}
	return
}

// Delete 轮播图删除
func (c *Controller) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return
}

// Update 轮播图更新
func (c *Controller) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, &model.RotationUpdateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		Id: req.Id,
	})
	return
}
