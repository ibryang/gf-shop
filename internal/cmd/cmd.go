package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop/internal/controller/admin"
	"shop/internal/controller/data"
	"shop/internal/controller/hello"
	"shop/internal/controller/login"
	"shop/internal/controller/permission"
	"shop/internal/controller/position"
	"shop/internal/controller/role"
	"shop/internal/controller/rotation"
	"shop/internal/service"
)

var (
	gfToken = &gtoken.GfToken{}
	Main    = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			gfToken = &gtoken.GfToken{
				ServerName:       "shop",
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  service.Login().LoginBeforeFunc,
				LoginAfterFunc:   service.Login().LoginAfterFunc,
				LogoutPath:       "/backend/logout",
				CacheMode:        2,
				GlobalMiddleware: false,
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthAfterFunc:    service.Login().AuthAfterFunc,
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				// gf默认的json返回格式
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().Ctx, service.Middleware().ResponseHandler)
				// gtoken
				// 鉴权认证
				group.Bind(admin.New().Create)
				//err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					hello.New(),
					rotation.New(),
					position.New(),
					login.New().Refresh,
					login.New().Logout,
					admin.New().Info,
					admin.New().List,
					admin.New().Delete,
					admin.New().Update,
					data.New(),
					role.New(),
					permission.New(),
				)
				//group.Middleware(service.Middleware().Auth)  // for jwt
				//group.ALLMap(map[string]interface{}{
				//	"/backend/admin/info": admin.New().Info,
				//})
			})
			s.Run()
			return nil
		},
	}
)
