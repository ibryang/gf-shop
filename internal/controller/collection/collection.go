package collection

import (
	"context"
	"shop/api/frontend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Create 收藏创建
func (c *Controller) Create(ctx context.Context, req *frontend.CollectionAddReq) (res *frontend.CollectionAddRes, err error) {
	output, err := service.Collection().Create(ctx, &model.CollectionCreateInput{
		ObjectId: req.ObjectId,
		Type:     req.Type,
	})
	if err != nil {
		return nil, err
	}
	res = &frontend.CollectionAddRes{
		Id: output.Id,
	}
	return
}

// Cancel 收藏取消
func (c *Controller) Cancel(ctx context.Context, req *frontend.CollectionCancelReq) (res *frontend.CollectionCancelRes, err error) {
	err = service.Collection().Cancel(ctx, &model.CollectionCancelInput{
		Id:       req.Id,
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	return
}

// List 列表
func (c *Controller) List(ctx context.Context, req *frontend.CollectionListReq) (res *frontend.CollectionListRes, err error) {
	output, err := service.Collection().List(ctx, &model.CollectionListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Type:     req.Type,
	})
	if err != nil {
		return nil, err
	}
	res = &frontend.CollectionListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}
