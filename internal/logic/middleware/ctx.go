package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/model"
	"shop/internal/service"
)

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 构建上下文
	customCtx := model.Context{
		Session: r.Session,
		Data:    g.Map{},
	}
	service.BizCtx().Init(r, &customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Username: user.Name,
			IsAdmin:  user.IsAdmin,
		}
	}
	r.Assigns(g.Map{
		"Context": &customCtx,
	})
	// 执行下一步
	r.Middleware.Next()
}
