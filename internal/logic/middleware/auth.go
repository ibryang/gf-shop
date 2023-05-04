package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/service"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
