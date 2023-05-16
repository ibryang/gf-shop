package frontend

import "github.com/gogf/gf/v2/frame/g"

type UserInfoBase struct {
	Name   string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Avatar string `json:"avatar" dc:"头像"`
	Sex    int    `json:"sex" dc:"1男 2女"`
	Sign   string `json:"sign" dc:"个性签名"`
	Status int    `json:"status" dc:"状态"`
}

type UserRegisterReq struct {
	g.Meta       `path:"/register" tags:"前台用户" method:"post" summary:"用户注册"`
	Name         string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password     string `json:"password" v:"password" dc:"密码"`
	Avatar       string `json:"avatar" dc:"头像"`
	Sex          int    `json:"sex" dc:"1男 2女"`
	Sign         string `json:"sign" dc:"个性签名"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}

type UserRegisterRes struct {
	Id int64 `json:"id"`
}

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"前台用户" method:"post" summary:"用户登录"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type UserLoginRes struct {
	Type   string      `json:"type"`
	Token  string      `json:"token"`
	Expire interface{} `json:"expire"`
	UserInfoBase
}

type UserInfoReq struct {
	g.Meta `path:"/frontend/user/info" tags:"前台用户" method:"get" summary:"用户信息"`
}

type UserInfoRes struct {
	Id int64 `json:"id" v:"required#用户id不能为空" dc:"用户id"`
	UserInfoBase
}

type UserUpdatePasswordReq struct {
	g.Meta       `path:"/frontend/user/update/password" tags:"前台用户" method:"post" summary:"用户修改密码"`
	Password     string `json:"password" v:"password" dc:"密码"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}

type UserUpdatePasswordRes struct {
	Id int64 `json:"id"`
}
