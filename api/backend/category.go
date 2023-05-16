package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type CategoryAddUpdateBaseReq struct {
	ParentId int    `json:"parent_id" dc:"parent_id" v:"" dc:"上级分类"`
	Name     string `json:"name" v:"required#分类名称不能为空" dc:"分类名称"`
	PicUrl   string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Level    int    `json:"level" dc:"等级 默认1级分类" d:"1"`
	Sort     int    `json:"sort" dc:"排序"`
}

type CategoryReq struct {
	g.Meta `path:"/backend/category/add" tags:"商品分类" method:"post" summary:"添加商品分类"`
	CategoryAddUpdateBaseReq
}

type CategoryRes struct {
	CategoryId int64 `json:"category_id" example:"1" dc:"category id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/backend/category/delete/{Id}" tags:"商品分类" method:"delete" summary:"删除商品分类"`
	Id     int `json:"id" v:"required|min:1#商品分类id不能为空|商品分类id最小为1" dc:"商品分类id"`
}

type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/backend/category/update/{Id}" tags:"商品分类" method:"put" summary:"更新商品分类"`
	Id     int `json:"id" v:"required|min:1#商品分类id不能为空|商品分类id最小为1" dc:"商品分类id"`
	CategoryAddUpdateBaseReq
}

type CategoryUpdateRes struct{}

type CategoryListReq struct {
	g.Meta `path:"/backend/category/list" tags:"商品分类" method:"get" summary:"商品分类列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type CategoryListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
