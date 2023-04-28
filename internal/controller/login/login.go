package login

import (
	"context"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c Controller) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	res = &backend.LoginRes{}
	err = service.Login().Login(ctx, &model.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res.Info = service.Session().GetUser(ctx)
	return
}
