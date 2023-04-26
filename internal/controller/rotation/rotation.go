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
