package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type PositionReq struct {
	g.Meta    `path:"/backend/position/add " tags:"手工位" method:"post" summary:"添加手工位"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link" description:"link" v:"required#商品链接不能为空" dc:"跳转链接"`
	Sort      int    `json:"sort" description:"sort" dc:"排序" d:"1"`
	GoodsName string `json:"goods_name" description:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   int    `json:"goods_id" description:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
}

type PositionRes struct {
	PositionId int64 `json:"position_id" example:"1" description:"position id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete/{Id}" tags:"手工位" method:"delete" summary:"删除手工位"`
	Id     int `json:"id" v:"required|min:1#手工位id不能为空|手工位id最小为1" dc:"手工位id"`
}

type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" tags:"手工位" method:"put" summary:"更新手工位"`
	Id        int    `json:"id" v:"required|min:1#手工位id不能为空|手工位id最小为1" dc:"手工位id"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link" description:"link" v:"required#商品链接不能为空" dc:"跳转链接"`
	Sort      int    `json:"sort" description:"sort" dc:"排序" d:"1"`
	GoodsName string `json:"goods_name" description:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   int    `json:"goods_id" description:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
}

type PositionUpdateRes struct{}

type PositionListReq struct {
	g.Meta `path:"/backend/position/list" tags:"手工位" method:"get" summary:"手工位列表"`
	Sort   int `json:"sort" description:"sort" dc:"排序"`
	common.PaginationReq
}

type PositionListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
