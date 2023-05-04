package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// LoginReq 登录
type LoginReq struct {
	g.Meta   `path:"/backend/login" method:"post" tags:"鉴权" summary:"登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// LoginRes 登录返回
type LoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// RefreshTokenReq 刷新token
type RefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"get" tags:"鉴权" summary:"刷新token"`
}

// RefreshTokenRes 刷新token返回
type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// LogoutReq 登出
type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"get" tags:"鉴权" summary:"登出"`
}
