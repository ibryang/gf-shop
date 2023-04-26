package rotation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
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

// Delete 轮播图删除
func (s *sRotation) Delete(ctx context.Context, id int) (err error) {
	return dao.Rotation.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// Unscoped() 取消软删除
		_, err = dao.Rotation.Ctx(ctx).Unscoped().Delete(dao.Rotation.Columns().Id, id)
		return err
	})
}
