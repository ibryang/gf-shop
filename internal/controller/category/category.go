package category

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

// List 商品分类列表
func (c *Controller) List(ctx context.Context, req *backend.CategoryListReq) (res *backend.CategoryListRes, err error) {
	output, err := service.Category().List(ctx, &model.CategoryListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.CategoryListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 商品分类创建
func (c *Controller) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	output, err := service.Category().Create(ctx, &model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Level:    req.Level,
			ParentId: req.ParentId,
			Sort:     req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.CategoryRes{
		CategoryId: output.CategoryId,
	}
	return
}

// Delete 商品分类删除
func (c *Controller) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

// Update 商品分类更新
func (c *Controller) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, &model.CategoryUpdateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Level:    req.Level,
			ParentId: req.ParentId,
			Sort:     req.Sort,
		},
		Id: req.Id,
	})
	return
}
