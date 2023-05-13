package permission

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

func (c *Controller) List(ctx context.Context, req *backend.PermissionListReq) (res *backend.PermissionListRes, err error) {
	list, err := service.Permission().List(ctx, &model.PermissionListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PermissionListRes{
		Total: list.Total,
		List:  list.List,
	}, nil
}

func (c *Controller) Create(ctx context.Context, req *backend.PermissionAddReq) (res *backend.PermissionAddRes, err error) {
	output, err := service.Permission().Create(ctx, &model.PermissionAddInput{
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.PermissionAddRes{
		PermissionId: output.PermissionId,
	}
	return
}

func (c *Controller) Delete(ctx context.Context, req *backend.PermissionDeleteReq) (res *backend.PermissionDeleteRes, err error) {
	err = service.Permission().Delete(ctx, req.Id)
	return &backend.PermissionDeleteRes{PermissionId: int64(req.Id)}, err
}

func (c *Controller) Update(ctx context.Context, req *backend.PermissionUpdateReq) (res *backend.PermissionUpdateRes, err error) {
	err = service.Permission().Update(ctx, &model.PermissionUpdateInput{
		Id: req.Id,
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	return &backend.PermissionUpdateRes{PermissionId: int64(req.Id)}, err
}
