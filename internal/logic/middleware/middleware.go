package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility/response"
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

// ResponseHandler 处理响应
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "", res)
	}

	return
}

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
			Username: user.Username,
			IsAdmin:  user.IsAdmin,
		}
	}
	r.Assigns(g.Map{
		"Context": &customCtx,
	})
	// 执行下一步
	r.Middleware.Next()
}
