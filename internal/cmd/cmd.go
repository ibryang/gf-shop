package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop/internal/consts"
	"shop/internal/controller/admin"
	"shop/internal/controller/article"
	"shop/internal/controller/category"
	"shop/internal/controller/coupon"
	"shop/internal/controller/data"
	"shop/internal/controller/file"
	"shop/internal/controller/goods"
	"shop/internal/controller/goods_options"
	"shop/internal/controller/hello"
	"shop/internal/controller/permission"
	"shop/internal/controller/position"
	"shop/internal/controller/role"
	"shop/internal/controller/rotation"
	"shop/internal/controller/user"
	"shop/internal/controller/user_coupon"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			GTokenBackend = RegisterBackendGToken()
			GTokenFrontend = RegisterFrontendGToken()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// SwaggerV2 格式的接口
				group.ALL("/swagger/v2", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})
				// gf默认的json返回格式
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().Ctx, service.Middleware().ResponseHandler)
				// gtoken
				// 鉴权认证
				group.Bind(admin.New().Create)
				//err := GTokenBackend.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					hello.New(),
					rotation.New(),
					position.New(),
					admin.New().Info,
					admin.New().List,
					admin.New().Delete,
					admin.New().Update,
					data.New(),
					role.New(),          // 角色
					permission.New(),    // 权限
					file.New(),          // 文件管理
					category.New(),      // 商品分类
					coupon.New(),        // 优惠券
					user_coupon.New(),   // 用户优惠券
					goods.New(),         // 商品
					goods_options.New(), // 商品规格
					article.New(),       // 文章
				)
				err = GTokenBackend.Middleware(ctx, group)
				group.Bind(article.New().Create)
				if err != nil {
					panic(err)
				}
				//group.Middleware(service.Middleware().Auth)  // for jwt
				//group.ALLMap(map[string]interface{}{
				//	"/backend/admin/info": admin.New().Info,
				//})
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Ctx, service.Middleware().ResponseHandler)
				group.Bind(
					user.New().Register,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err = GTokenFrontend.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						user.New().Info,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
