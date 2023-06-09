// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IPermission interface {
		List(ctx context.Context, in *model.PermissionListInput) (out *model.PermissionItemListOutput, err error)
		Create(ctx context.Context, in *model.PermissionAddInput) (out *model.PermissionAddOutput, err error)
		Delete(ctx context.Context, id int) (err error)
		Update(ctx context.Context, in *model.PermissionUpdateInput) (err error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
