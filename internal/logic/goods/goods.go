package goods

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sGoods struct{}

func init() {
	// 注册路由
	service.RegisterGoods(New())
}
func New() *sGoods {
	return &sGoods{}
}

// List 商品列表
func (s *sGoods) List(ctx context.Context, in *model.GoodsListInput) (out *model.GoodsListItemOutput, err error) {
	var m = dao.GoodsInfo.Ctx(ctx)

	out = &model.GoodsListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.GoodsInfo.Columns().Id)
	var list []*model.GoodsListItem
	if err = listModel.Scan(&list); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return out, nil
	}

	out.Total, err = m.Count()
	if err != nil {
		return nil, err
	}
	out.List = list
	return out, nil
}

// Create 商品创建
func (s *sGoods) Create(ctx context.Context, in *model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	id, err := dao.GoodsInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.GoodsCreateOutput{}, err
	}
	return model.GoodsCreateOutput{Id: id}, nil
}

// Delete 商品删除
func (s *sGoods) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.GoodsInfo.Ctx(ctx).Delete(dao.GoodsInfo.Columns().Id, id)
	return err
}

// Update 商品更新
func (s *sGoods) Update(ctx context.Context, in *model.GoodsUpdateInput) (err error) {
	_, err = dao.GoodsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsInfo.Columns().Id).
		Where(dao.GoodsInfo.Columns().Id, in.Id).
		Update()
	return err
}
