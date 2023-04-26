package rotation

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
