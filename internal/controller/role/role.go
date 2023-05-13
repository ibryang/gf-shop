package role

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

func (c *Controller) List(ctx context.Context, req *backend.RoleListReq) (res *backend.RoleListRes, err error) {
	list, err := service.Role().List(ctx, &model.RoleListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleListRes{
		Total: list.Total,
		List:  list.List,
	}, nil
}

func (c *Controller) Create(ctx context.Context, req *backend.RoleAddReq) (res *backend.RoleAddRes, err error) {
	output, err := service.Role().Create(ctx, &model.RoleAddInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.RoleAddRes{
		RoleId: output.RoleId,
	}
	return
}

func (c *Controller) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return &backend.RoleDeleteRes{RoleId: int64(req.Id)}, err
}

func (c *Controller) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, &model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return &backend.RoleUpdateRes{RoleId: int64(req.Id)}, err
}
