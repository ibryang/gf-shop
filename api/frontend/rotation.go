package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type RotationListReq struct {
	g.Meta `path:"/frontend/rotation/list" tags:"轮播图" method:"get" summary:"轮播图列表"`
	Sort   int `json:"sort" description:"sort" dc:"排序"`
	common.PaginationReq
}

type RotationListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
