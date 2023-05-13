package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sRole struct{}

func init() {
	// 注册路由
	service.RegisterRole(New())
}
func New() *sRole {
	return &sRole{}
}

// List 用户列表
func (s *sRole) List(ctx context.Context, in *model.RoleListInput) (out *model.RoleItemListOutput, err error) {
	out = &model.RoleItemListOutput{}
	var m = dao.RoleInfo.Ctx(ctx)
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.RoleInfo.Columns().Id)
	var list []*model.RoleItemOutput
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

// Create 添加角色
func (s *sRole) Create(ctx context.Context, in *model.RoleAddInput) (out *model.RoleAddOutput, err error) {
	lastInsertId, err := dao.RoleInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}

	return &model.RoleAddOutput{RoleId: lastInsertId}, nil
}

// AddPermission 添加权限
func (s *sRole) AddPermission(ctx context.Context, in *model.RoleAddPermissionInput) (out *model.RoleAddPermissionOutput, err error) {
	id, err := dao.RolePermissionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}
	return &model.RoleAddPermissionOutput{Id: id}, nil
}

// Delete 角色删除
func (s *sRole) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.RoleInfo.Ctx(ctx).Unscoped().Delete(dao.RoleInfo.Columns().Id, id)
	return
}

// DeletePermission 删除权限
func (s *sRole) DeletePermission(ctx context.Context, in *model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).Where(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	return err
}

// Update 用户更新
func (s *sRole) Update(ctx context.Context, in *model.RoleUpdateInput) (err error) {
	_, err = dao.RoleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}
