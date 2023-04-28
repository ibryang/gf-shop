package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility"
)

type sUser struct{}

func init() {
	// 注册路由
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

// List 用户列表
func (s *sUser) List(ctx context.Context, in *model.UserListInput) (out *model.UserListItemOutput, err error) {
	out = &model.UserListItemOutput{}
	var m = dao.User.Ctx(ctx)
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.User.Columns().Sort)
	var list []*model.UserListItem
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

// IsExist 查询用户是否存在
func (s *sUser) IsExist(ctx context.Context, username string) bool {
	var user *model.UserListItem
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Scan(&user)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

// Create 用户创建
func (s *sUser) Create(ctx context.Context, in *model.UserCreateInput) (out model.UserCreateOutput, err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Username) {
		return model.UserCreateOutput{}, fmt.Errorf("用户已存在")
	}
	// 加密密码
	salt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, salt)
	in.Salt = salt
	id, err := dao.User.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.UserCreateOutput{}, err
	}
	return model.UserCreateOutput{RequestId: id}, nil
}

// Delete 用户删除
func (s *sUser) Delete(ctx context.Context, id int) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Delete(dao.User.Columns().Id, id)
		return err
	})
}

// Update 用户更新
func (s *sUser) Update(ctx context.Context, in *model.UserUpdateInput) (err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Username) {
		return fmt.Errorf("用户已存在")
	}
	// 判断是否修改了密码
	if in.Password != "" {
		// 加密密码
		salt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, salt)
		in.Salt = salt
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.User.Columns().Id).
			Where(dao.User.Columns().Id, in.Id).
			Update()
		return err
	})
}
