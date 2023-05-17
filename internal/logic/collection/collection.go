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

func (s *sCollection) List(ctx context.Context, in *model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	var m = dao.CollectionInfo.Ctx(ctx)
	out = &model.CollectionListOutput{}
	listModel := m.Page(in.Page, in.PageSize)

	//var userId = ctx.Value(consts.ContextUserId).(int)
	//listModel = listModel.Where(g.Map{
	//	dao.CollectionInfo.Columns().UserId: userId,
	//})
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}
	var list []*model.CollectionItem
	//if err = listModel.WithAll().Scan(&list); err != nil {  // WithAll() 查询所有的关联数据
	if in.Type == 1 {
		err = listModel.With(model.GoodsItem{}).Scan(&list)
	} else if in.Type == 2 {
		err = listModel.With(model.ArticleItem{}).Scan(&list)
	}
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return out, nil
	}

	out.Total, err = listModel.Count()
	if err != nil {
		return nil, err
	}
	out.List = list
	return out, nil
}
