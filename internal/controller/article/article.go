package article

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

// List 文章列表
func (c *Controller) List(ctx context.Context, req *backend.ArticleListReq) (res *backend.ArticleListRes, err error) {
	output, err := service.Article().List(ctx, &model.ArticleListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.ArticleListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 文章创建
func (c *Controller) Create(ctx context.Context, req *backend.ArticleAddReq) (res *backend.ArticleRes, err error) {
	var data = model.ArticleCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	output, err := service.Article().Create(ctx, &data)
	if err != nil {
		return nil, err
	}
	res = &backend.ArticleRes{
		Id: output.Id,
	}
	return
}

// Delete 文章删除
func (c *Controller) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

// Update 文章更新
func (c *Controller) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	var data = model.ArticleUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Article().Update(ctx, &data)
	if err != nil {
		return nil, err
	}
	return
}
