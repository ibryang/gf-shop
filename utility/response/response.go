package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type JsonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var resData interface{}
	if len(data) > 0 {
		resData = data[0]
	} else {
		resData = nil
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    resData,
	})
}

func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}
