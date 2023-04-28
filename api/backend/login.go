package backend

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 登录
type LoginReq struct {
	g.Meta   `path:"/backend/login" method:"post" tags:"鉴权" summary:"登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// LoginRes 登录返回
type LoginRes struct {
	Info interface{} `json:"info"`
}
