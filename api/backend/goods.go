package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type GoodsAddUpdateBaseReq struct {
	PicUrl           string `json:"pic_url" dc:"图片"`
	Name             string `json:"name" v:"required#商品名称不能为空" dc:"商品名称"`
	Price            int    `json:"price" v:"required#价格不能为空" dc:"价格 单位分"`
	Level1CategoryId int    `json:"level1_category_id" dc:"1级分类id"`
	Level2CategoryId int    `json:"level2_category_id" dc:"2级分类id"`
	Level3CategoryId int    `json:"level3_category_id" dc:"3级分类id"`
	Brand            string `json:"brand" v:"max-length:30#品牌名称最大30个字符" dc:"品牌"`
	Stock            int    `json:"stock" dc:"库存"`
	Sale             int    `json:"sale" dc:"销量"`
	Tags             string `json:"tags" dc:"标签"`
	DetailInfo       string `json:"detail_info" dc:"商品详情"`
}

type GoodsAddReq struct {
	g.Meta `path:"/backend/goods/add" tags:"商品" method:"post" summary:"添加商品"`
	GoodsAddUpdateBaseReq
}

type GoodsRes struct {
	Id int64 `json:"id" example:"1" dc:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/backend/goods/delete/{Id}" tags:"商品" method:"delete" summary:"删除商品"`
	Id     int `json:"id" v:"required|min:1#商品id不能为空|商品id最小为1" dc:"商品id"`
}

type GoodsDeleteRes struct{}

type GoodsUpdateReq struct {
	g.Meta `path:"/backend/goods/update/{Id}" tags:"商品" method:"put" summary:"更新商品"`
	Id     int `json:"id" v:"required|min:1#商品id不能为空|商品id最小为1" dc:"商品id"`
	GoodsAddUpdateBaseReq
}

type GoodsUpdateRes struct{}

type GoodsListReq struct {
	g.Meta `path:"/backend/goods/list" tags:"商品" method:"get" summary:"商品列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type GoodsListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
