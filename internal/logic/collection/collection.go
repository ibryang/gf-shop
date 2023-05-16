package coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sCollection struct{}

func init() {
	// 注册路由
	service.RegisterCollection(New())
}
func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) Create(ctx context.Context, in *model.CollectionCreateInput) (out *model.CollectionCreateOutput, err error) {
	userId := ctx.Value(consts.ContextUserId).(int)
	in.UserId = userId
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}
	return &model.CollectionCreateOutput{Id: int(id)}, nil
}

func (s *sCollection) Cancel(ctx context.Context, in *model.CollectionCancelInput) (err error) {
	userId := ctx.Value(consts.ContextUserId).(int)
	if in.Id != 0 {
		_, err = dao.CollectionInfo.Ctx(ctx).Where(g.Map{
			dao.CollectionInfo.Columns().UserId: userId,
			dao.CollectionInfo.Columns().Id:     in.Id,
		},
		).Delete()
	} else {
		_, err = dao.CollectionInfo.Ctx(ctx).Where(g.Map{
			dao.CollectionInfo.Columns().UserId:   userId,
			dao.CollectionInfo.Columns().Type:     in.Type,
			dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		}).Delete()
	}
	return err
}
