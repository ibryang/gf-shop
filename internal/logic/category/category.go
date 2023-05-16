package category

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sCategory struct{}

func init() {
	// 注册路由
	service.RegisterCategory(New())
}
func New() *sCategory {
	return &sCategory{}
}

// List 商品分类列表
func (s *sCategory) List(ctx context.Context, in *model.CategoryListInput) (out *model.CategoryListItemOutput, err error) {
	var m = dao.CategoryInfo.Ctx(ctx)

	out = &model.CategoryListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	var list []*model.CategoryListItem
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

// Create 商品分类创建
func (s *sCategory) Create(ctx context.Context, in *model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	id, err := dao.CategoryInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.CategoryCreateOutput{}, err
	}
	return model.CategoryCreateOutput{CategoryId: id}, nil
}

// Delete 商品分类删除
func (s *sCategory) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.CategoryInfo.Ctx(ctx).Delete(dao.CategoryInfo.Columns().Id, id)
	return err
}

// Update 商品分类更新
func (s *sCategory) Update(ctx context.Context, in *model.CategoryUpdateInput) (err error) {
	_, err = dao.CategoryInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return err
}
