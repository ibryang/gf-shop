package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
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

// IsExist 查询用户是否存在
func (s *sUser) IsExist(ctx context.Context, username string) bool {
	var user *entity.UserInfo
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

// Register 用户注册
func (s *sUser) Register(ctx context.Context, in *model.UserRegisterInput) (out model.UserRegisterOutput, err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Name) {
		return model.UserRegisterOutput{}, fmt.Errorf(consts.ErrUserIsExistMsg)
	}
	// 加密密码
	salt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, salt)
	in.UserSalt = salt
	id, err := dao.UserInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.UserRegisterOutput{}, err
	}
	return model.UserRegisterOutput{Id: id}, nil
}
