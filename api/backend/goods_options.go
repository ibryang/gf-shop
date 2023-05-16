package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type GoodsOptionsAddUpdateBaseReq struct {
	GoodsId int    `json:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
	Name    string `json:"name" v:"required#商品规格名称不能为空" dc:"商品规格名称"`
	Price   int    `json:"price" v:"required#价格不能为空" dc:"价格 单位分"`
	Brand   string `json:"brand" v:"max-length:30#品牌名称最大30个字符" dc:"品牌"`
	PicUrl  string `json:"pic_url" dc:"图片"`
	Stock   int    `json:"stock" dc:"库存"`
}

type GoodsOptionsAddReq struct {
	g.Meta `path:"/backend/goods/options/add" tags:"商品规格" method:"post" summary:"添加商品规格"`
	GoodsOptionsAddUpdateBaseReq
}

type GoodsOptionsRes struct {
	Id int64 `json:"id" example:"1" dc:"id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/backend/goods/options/delete/{Id}" tags:"商品规格" method:"delete" summary:"删除商品规格"`
	Id     int `json:"id" v:"required|min:1#商品规格id不能为空|商品规格id最小为1" dc:"商品规格id"`
}

type GoodsOptionsDeleteRes struct{}

type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/backend/goods/options/update/{Id}" tags:"商品规格" method:"put" summary:"更新商品规格"`
	Id     int `json:"id" v:"required|min:1#商品规格id不能为空|商品规格id最小为1" dc:"商品规格id"`
	GoodsOptionsAddUpdateBaseReq
}

type GoodsOptionsUpdateRes struct{}

type GoodsOptionsListReq struct {
	g.Meta `path:"/backend/goods/options/list" tags:"商品规格" method:"get" summary:"商品规格列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type GoodsOptionsListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
