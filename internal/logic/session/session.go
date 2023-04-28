package session

import (
	"context"
	"shop/internal/model/entity"
	"shop/internal/service"
)

type sSession struct{}

const (
	sessionUserKey         = "SessionUserKey"
	sessionKeyLoginReferer = "SessionKeyReferer"
	sessionKeyNotice       = "SessionKeyNotice"
)

func New() *sSession {
	return &sSession{}
}

func init() {
	// 注册路由
	service.RegisterSession(New())
}

// SetUser 将用户信息设置到上下文中, 供请求使用
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.BizCtx().Get(ctx).Session.Set(sessionUserKey, user)
}

// GetUser 获取上下文用户
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionUserKey)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.User{}
}

// RemoveUser 移除上下文用户
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionUserKey)
	}
	return nil
}
