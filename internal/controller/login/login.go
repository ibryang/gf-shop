package login

import (
	"context"
	"shop/api/backend"
	"shop/internal/service"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

//func (c Controller) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
//	res = &backend.LoginRes{}
//	err = service.Login().Login(ctx, &model.LoginInput{
//		Username: req.Username,
//		Password: req.Password,
//	})
//	if err != nil {
//		return nil, err
//	}
//	//res.Token = service.Session().Token
//	return
//}

// Login 登录
func (c Controller) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	res = &backend.LoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

// Refresh 刷新token
func (c Controller) Refresh(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

// Logout 退出
func (c Controller) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LoginRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
