package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop/internal/controller/hello"
	"shop/internal/controller/login"
	"shop/internal/controller/position"
	"shop/internal/controller/rotation"
	"shop/internal/controller/user"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().Ctx, service.Middleware().ResponseHandler)
				group.Bind(
					hello.New(),
					rotation.New(),
					position.New(),
					user.New(),
					login.New(),
				)
				group.Middleware(service.Middleware().Auth)
				group.ALLMap(map[string]interface{}{
					"/backend/user/info": user.New().Info,
				})
			})
			s.Run()
			return nil
		},
	}
)
