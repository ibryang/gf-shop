package rotation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sPosition struct{}

func init() {
	// 注册路由
	service.RegisterPosition(New())
}
func New() *sPosition {
	return &sPosition{}
}

// List 手工位列表
func (s *sPosition) List(ctx context.Context, in *model.PositionListInput) (out *model.PositionListItemOutput, err error) {
	out = &model.PositionListItemOutput{}
	var m = dao.Position.Ctx(ctx)

	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.Position.Columns().Sort)
	var list []*model.PositionListItem
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

// Create 手工位创建
func (s *sPosition) Create(ctx context.Context, in *model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	id, err := dao.Position.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.PositionCreateOutput{}, err
	}
	return model.PositionCreateOutput{RequestId: id}, nil
}

// Delete 手工位删除
func (s *sPosition) Delete(ctx context.Context, id int) (err error) {
	return dao.Position.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Position.Ctx(ctx).Delete(dao.Position.Columns().Id, id)
		return err
	})
}

// Update 手工位更新
func (s *sPosition) Update(ctx context.Context, in *model.PositionUpdateInput) (err error) {
	return dao.Position.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许html代码
		if err = ghtml.SpecialCharsMapOrStruct(in.PicUrl); err != nil {
			return err
		}
		_, err = dao.Position.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Position.Columns().Id).
			Where(dao.Position.Columns().Id, in.Id).
			Update()
		return err
	})
}
