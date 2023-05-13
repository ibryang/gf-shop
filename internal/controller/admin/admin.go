package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/backend"
	"shop/internal/consts"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Create 用户创建
func (c *Controller) Create(ctx context.Context, req *backend.AdminAddReq) (res *backend.AdminAddRes, err error) {
	output, err := service.Admin().Create(ctx, &model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Username,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &backend.AdminAddRes{
		AdminId: output.RequestId,
	}
	return
}

// List 用户列表
func (c *Controller) List(ctx context.Context, req *backend.AdminListReq) (res *backend.AdminListRes, err error) {
	output, err := service.Admin().List(ctx, &model.AdminListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &backend.AdminListRes{
		List:  output.List,
		Total: output.Total,
	}
	return
}

// Delete 用户删除
func (c *Controller) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

// Update 用户更新
func (c *Controller) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, &model.AdminUpdateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Username,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
		},
		Id: req.Id,
	})
	return
}

// Info 用户详情
// for jwt
//func (c *Controller) Info(ctx context.Context, req *backend.AdminInfoReq) (res *backend.AdminInfoRes, err error) {
//	return &backend.AdminInfoRes{
//		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
//		IdentityKey: service.Auth().IdentityKey,
//		Payload:     service.Auth().GetPayload(ctx),
//	}, nil
//}

// Info 用户详情
// for gtoken
func (c *Controller) Info(ctx context.Context, req *backend.AdminInfoReq) (res *backend.AdminInfoRes, err error) {
	res = &backend.AdminInfoRes{
		Id:       gconv.Int(ctx.Value(consts.ContextUserId)),
		Username: gconv.String(ctx.Value(consts.ContextUsername)),
		IsAdmin:  gconv.Int(ctx.Value(consts.ContextUserIsAdmin)),
		RoleIds:  gconv.Ints(ctx.Value(consts.ContextUserRoleIds)),
	}
	return
}
