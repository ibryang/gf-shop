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
	IRole interface {
		List(ctx context.Context, in *model.RoleListInput) (out *model.RoleItemListOutput, err error)
		Create(ctx context.Context, in *model.RoleAddInput) (out *model.RoleAddOutput, err error)
		Delete(ctx context.Context, id int) (err error)
		Update(ctx context.Context, in *model.RoleUpdateInput) (err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}