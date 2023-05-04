package middleware

import (
	"shop/internal/service"
)

type sMiddleware struct {
	LoginUrl string
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/backend/login",
	}
}

func init() {
	// 注册路由
	service.RegisterMiddleware(New())
}
