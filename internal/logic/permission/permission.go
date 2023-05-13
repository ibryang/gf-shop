package permission

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sPermission struct{}

func init() {
	// 注册路由
	service.RegisterPermission(New())
}
func New() *sPermission {
	return &sPermission{}
}

// List 用户列表
func (s *sPermission) List(ctx context.Context, in *model.PermissionListInput) (out *model.PermissionItemListOutput, err error) {
	out = &model.PermissionItemListOutput{}
	var m = dao.PermissionInfo.Ctx(ctx)
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.PermissionInfo.Columns().Id)
	var list []*model.PermissionItemOutput
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

// Create 权限添加
func (s *sPermission) Create(ctx context.Context, in *model.PermissionAddInput) (out *model.PermissionAddOutput, err error) {
	lastInsertId, err := dao.PermissionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}

	return &model.PermissionAddOutput{PermissionId: lastInsertId}, nil
}

// Delete 权限删除
func (s *sPermission) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PermissionInfo.Ctx(ctx).Unscoped().Delete(dao.PermissionInfo.Columns().Id, id)
	return
}

// Update 用户更新
func (s *sPermission) Update(ctx context.Context, in *model.PermissionUpdateInput) (err error) {
	_, err = dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err
}
