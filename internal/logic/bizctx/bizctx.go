package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/consts"
	"shop/internal/model"
	"shop/internal/service"
)

type sBizCtx struct{}

func New() *sBizCtx {
	return &sBizCtx{}
}

func init() {
	// 注册路由
	service.RegisterBizCtx(New())
}

// Init 初始化上下文对象到请求上下文,以便后续的请求流程中可以修改
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获取上下文
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将用户信息设置到上下文中, 供请求使用
func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// SetData 将数据设置到上下文中, 供请求使用
func (s *sBizCtx) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
