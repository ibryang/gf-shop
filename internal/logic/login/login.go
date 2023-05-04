package login

import (
	"context"
	"fmt"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
)

type sLogin struct{}

func init() {
	// 注册路由
	service.RegisterLogin(New())
}
func New() *sLogin {
	return &sLogin{}
}

func (s *sLogin) Login(ctx context.Context, in *model.LoginInput) (err error) {
	var user = new(entity.User)
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		return err
	}
	encryptPassword := utility.EncryptPassword(in.Password, user.Salt)
	if user.Password != encryptPassword {
		return fmt.Errorf("用户名或密码错误")
	}
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	})

	return
}
