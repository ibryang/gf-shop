// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUser interface {
		IsExist(ctx context.Context, username string) bool
		Register(ctx context.Context, in *model.UserRegisterInput) (out model.UserRegisterOutput, err error)
		LoginBeforeFunc(r *ghttp.Request) (string, interface{})
		LoginAfterFunc(r *ghttp.Request, resp gtoken.Resp)
		AuthAfterFunc(r *ghttp.Request, resp gtoken.Resp)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}