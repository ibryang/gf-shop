package frontend

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	g.Meta       `path:"/register" tags:"前台注册" method:"post" summary:"用户注册"`
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
