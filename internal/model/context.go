package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // 当前Session的管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 上下文数据, 格式K:V, 不固定
}

type ContextUser struct {
	Id       int
	Username string
	IsAdmin  int
}
