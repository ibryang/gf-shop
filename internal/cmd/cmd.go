package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop/internal/consts"
	"shop/internal/controller/admin"
	"shop/internal/controller/category"
	"shop/internal/controller/coupon"
	"shop/internal/controller/data"
	"shop/internal/controller/file"
	"shop/internal/controller/goods"
	"shop/internal/controller/hello"
	"shop/internal/controller/login"
	"shop/internal/controller/permission"
	"shop/internal/controller/position"
	"shop/internal/controller/role"
	"shop/internal/controller/rotation"
	"shop/internal/controller/userCoupon"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			GToken = RegisterGToken()
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
				//err := GToken.Middleware(ctx, group)
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
					role.New(),       // 角色
					permission.New(), // 权限
					file.New(),       // 文件管理
					category.New(),   // 商品分类
					coupon.New(),     // 优惠券
					userCoupon.New(), // 用户优惠券
					goods.New(),      // 商品
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
