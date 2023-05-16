package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/frontend"
	"shop/internal/consts"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Register 用户注册
func (c *Controller) Register(ctx context.Context, req *frontend.UserRegisterReq) (res *frontend.UserRegisterRes, err error) {
	var data = model.UserRegisterInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, &data)
	if err != nil {
		return nil, err
	}
	res = &frontend.UserRegisterRes{
		Id: out.Id,
	}
	return
}

// Info 用户信息
func (c *Controller) Info(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	res = &frontend.UserInfoRes{
		Id: gconv.Int64(ctx.Value(consts.ContextAdminId)),
		UserInfoBase: frontend.UserInfoBase{
			Name:   gconv.String(ctx.Value(consts.ContextUserName)),
			Avatar: gconv.String(ctx.Value(consts.ContextUserAvatar)),
			Sex:    gconv.Int(ctx.Value(consts.ContextUserSex)),
			Sign:   gconv.String(ctx.Value(consts.ContextUserSign)),
			Status: gconv.Int(ctx.Value(consts.ContextUserStatus)),
		},
	}
	return
}
