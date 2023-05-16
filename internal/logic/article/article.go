package goods

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sArticle struct{}

func init() {
	// 注册路由
	service.RegisterArticle(New())
}
func New() *sArticle {
	return &sArticle{}
}

// List 文章列表
func (s *sArticle) List(ctx context.Context, in *model.ArticleListInput) (out *model.ArticleListItemOutput, err error) {
	var m = dao.ArticleInfo.Ctx(ctx)

	out = &model.ArticleListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.ArticleInfo.Columns().Id)
	var list []*model.ArticleListItem
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

// Create 文章创建
func (s *sArticle) Create(ctx context.Context, in *model.ArticleCreateInput) (out model.ArticleCreateOutput, err error) {
	id, err := dao.ArticleInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.ArticleCreateOutput{}, err
	}
	return model.ArticleCreateOutput{Id: id}, nil
}

// Delete 文章删除
func (s *sArticle) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.ArticleInfo.Ctx(ctx).Delete(dao.ArticleInfo.Columns().Id, id)
	return err
}

// Update 文章更新
func (s *sArticle) Update(ctx context.Context, in *model.ArticleUpdateInput) (err error) {
	_, err = dao.ArticleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.ArticleInfo.Columns().Id).
		Where(dao.ArticleInfo.Columns().Id, in.Id).
		Update()
	return err
}
