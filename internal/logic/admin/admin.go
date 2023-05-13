package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
)

type sAdmin struct{}

func init() {
	// 注册路由
	service.RegisterAdmin(New())
}
func New() *sAdmin {
	return &sAdmin{}
}

// List 用户列表
func (s *sAdmin) List(ctx context.Context, in *model.AdminListInput) (out *model.AdminItemListOutput, err error) {
	out = &model.AdminItemListOutput{}
	var m = dao.AdminInfo.Ctx(ctx)
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.AdminInfo.Columns().Id)
	var list []*model.AdminItemOutput
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
func (s *sAdmin) IsExist(ctx context.Context, username string) bool {
	var user *model.AdminItemOutput
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

// Create 用户创建
func (s *sAdmin) Create(ctx context.Context, in *model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Name) {
		return model.AdminCreateOutput{}, fmt.Errorf("用户已存在")
	}
	// 加密密码
	salt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, salt)
	in.UserSalt = salt
	id, err := dao.AdminInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AdminCreateOutput{}, err
	}
	return model.AdminCreateOutput{RequestId: id}, nil
}

// Delete 用户删除
func (s *sAdmin) Delete(ctx context.Context, id int) (err error) {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.AdminInfo.Ctx(ctx).Delete(dao.AdminInfo.Columns().Id, id)
		return err
	})
}

// Update 用户更新
func (s *sAdmin) Update(ctx context.Context, in *model.AdminUpdateInput) (err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Name) {
		return fmt.Errorf("用户已存在")
	}
	// 判断是否修改了密码
	if in.Password != "" {
		// 加密密码
		salt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, salt)
		in.UserSalt = salt
	}
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetUserByUsernamePassword 通过用户名和密码获取用户信息
func (s *sAdmin) GetUserByUsernamePassword(ctx context.Context, in model.LoginInput) (out map[string]interface{}) {
	var (
		user *entity.AdminInfo
	)
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Username).Scan(&user)
	if err != nil {
		return nil
	}
	if user == nil {
		return nil
	}
	if user.Password != utility.EncryptPassword(in.Password, user.UserSalt) {
		return nil
	}
	return g.Map{
		"id":       user.Id,
		"username": user.Name,
		"is_admin": user.IsAdmin,
		"role_ids": user.RoleIds,
	}
}
