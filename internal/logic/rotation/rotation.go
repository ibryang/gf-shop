package rotation

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sRotation struct{}

func init() {
	// 注册路由
	service.RegisterRotation(New())
}
func New() *sRotation {
	return &sRotation{}
}

// Create 轮播图创建
func (s *sRotation) Create(ctx context.Context, in *model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	id, err := dao.Rotation.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.RotationCreateOutput{}, err
	}
	return model.RotationCreateOutput{RequestId: id}, nil
}
