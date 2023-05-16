package cmd

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/service"
)

var GTokenBackend *gtoken.GfToken
var GTokenFrontend *gtoken.GfToken

func RegisterBackendGToken() *gtoken.GfToken {
	GTokenBackend = &gtoken.GfToken{
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
	return GTokenBackend
}

func RegisterFrontendGToken() *gtoken.GfToken {
	GTokenFrontend = &gtoken.GfToken{
		ServerName:      "shop",
		LoginPath:       "/login",
		LoginBeforeFunc: service.User().LoginBeforeFunc,
		LoginAfterFunc:  service.User().LoginAfterFunc,
		LogoutPath:      "/logout",
		CacheMode:       2,
		MultiLogin:      false,
		AuthAfterFunc:   service.User().AuthAfterFunc,
	}
	return GTokenFrontend
}
