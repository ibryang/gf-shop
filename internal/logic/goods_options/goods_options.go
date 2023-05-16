package goods_options

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sGoodsOptions struct{}

func init() {
	service.RegisterGoodsOptions(New())
}
func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

// List 商品规格列表
func (s *sGoodsOptions) List(ctx context.Context, in *model.GoodsOptionsListInput) (out *model.GoodsOptionsListItemOutput, err error) {
	var m = dao.GoodsOptionsInfo.Ctx(ctx)

	out = &model.GoodsOptionsListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.GoodsOptionsInfo.Columns().Id)
	var list []*model.GoodsOptionsListItem
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

// Create 商品规格创建
func (s *sGoodsOptions) Create(ctx context.Context, in *model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error) {
	id, err := dao.GoodsOptionsInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.GoodsOptionsCreateOutput{}, err
	}
	return model.GoodsOptionsCreateOutput{Id: id}, nil
}

// Delete 商品规格删除
func (s *sGoodsOptions) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.GoodsOptionsInfo.Ctx(ctx).Delete(dao.GoodsOptionsInfo.Columns().Id, id)
	return err
}

// Update 商品规格更新
func (s *sGoodsOptions) Update(ctx context.Context, in *model.GoodsOptionsUpdateInput) (err error) {
	_, err = dao.GoodsOptionsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsOptionsInfo.Columns().Id).
		Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).
		Update()
	return err
}
