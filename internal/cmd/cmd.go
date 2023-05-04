package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
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
			gfToken := &gtoken.GfToken{
				ServerName:       "shop",
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  service.Login().LoginBeforeFunc,
				LoginAfterFunc:   service.Login().LoginAfterFunc,
				LogoutPath:       "/backend/logout",
				CacheMode:        2,
				GlobalMiddleware: false,
				AuthPaths:        g.SliceStr{"/backend/user/info"},
				AuthAfterFunc:    service.Login().AuthAfterFunc,
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				// gf默认的json返回格式
				//group.Middleware(ghttp.MiddlewareHandlerResponse)

				// gtoken
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Middleware(service.Middleware().Ctx, service.Middleware().ResponseHandler)
				group.Bind(
					hello.New(),
					rotation.New(),
					position.New(),
					user.New(),
					login.New().Refresh,
					login.New().Logout,
				)
				//group.Middleware(service.Middleware().Auth)  // for jwt
				//group.ALLMap(map[string]interface{}{
				//	"/backend/user/info": user.New().Info,
				//})
			})
			s.Run()
			return nil
		},
	}
)
