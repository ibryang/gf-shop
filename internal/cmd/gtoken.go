package cmd

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/service"
)

var GToken *gtoken.GfToken

func RegisterGToken() *gtoken.GfToken {
	GToken = &gtoken.GfToken{
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
	return GToken
}
