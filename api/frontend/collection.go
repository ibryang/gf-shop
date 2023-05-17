package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type CollectionAddReq struct {
	g.Meta   `path:"/frontend/collection/add" tags:"前台收藏" method:"post" summary:"添加收藏"`
	ObjectId int `json:"object_id" v:"required#对象id不能为空" dc:"对象id"`
	Type     int `json:"type" v:"in:1,2" dc:"类型: 1商品, 2文章"`
}

type CollectionAddRes struct {
	Id int `json:"id"`
}

type CollectionCancelReq struct {
	g.Meta   `path:"/frontend/collection/cancel" tags:"前台收藏" method:"post" summary:"取消收藏"`
	Id       int `json:"id" dc:"收藏id"`
	ObjectId int `json:"object_id" dc:"对象id"`
	Type     int `json:"type" dc:"类型: 1商品, 2文章"`
}

type CollectionCancelRes struct {
}

type CollectionListReq struct {
	g.Meta `path:"/frontend/collection/list" tags:"前台收藏" method:"get" summary:"收藏列表"`
	common.PaginationReq
}

type CollectionListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
