package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// List 用户列表
func (c *Controller) List(ctx context.Context, req *backend.UserListReq) (res *backend.UserListRes, err error) {
	output, err := service.User().List(ctx, &model.UserListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.UserListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Create 用户创建
func (c *Controller) Create(ctx context.Context, req *backend.UserReq) (res *backend.UserRes, err error) {
	output, err := service.User().Create(ctx, &model.UserCreateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Username: req.Username,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
			Sort:     req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.UserRes{
		UserId: output.RequestId,
	}
	return
}

// Delete 用户删除
func (c *Controller) Delete(ctx context.Context, req *backend.UserDeleteReq) (res *backend.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	return
}

// Update 用户更新
func (c *Controller) Update(ctx context.Context, req *backend.UserUpdateReq) (res *backend.UserUpdateRes, err error) {
	err = service.User().Update(ctx, &model.UserUpdateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Username: req.Username,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
			Sort:     req.Sort,
		},
		Id: req.Id,
	})
	return
}

// Info 用户详情
// for jwt
//func (c *Controller) Info(ctx context.Context, req *backend.UserInfoReq) (res *backend.UserInfoRes, err error) {
//	return &backend.UserInfoRes{
//		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
//		IdentityKey: service.Auth().IdentityKey,
//		Payload:     service.Auth().GetPayload(ctx),
//	}, nil
//}

// Info 用户详情
// for gtoken
func (c *Controller) Info(ctx context.Context, req *backend.UserInfoReq) (res *backend.UserInfoRes, err error) {
	fmt.Println(gconv.Int(ctx.Value("id")))
	res = &backend.UserInfoRes{
		Id:       gconv.Int(ctx.Value("id")),
		Username: gconv.String(ctx.Value("username")),
		IsAdmin:  gconv.Int(ctx.Value("isAdmin")),
		RoleIds:  gconv.Ints(ctx.Value("roleIds")),
	}
	fmt.Println(res)
	return

}
