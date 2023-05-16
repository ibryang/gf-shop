package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"轮播图" method:"post" summary:"添加轮播图"`
	PicUrl string `json:"pic_url" v:"required:图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link" description:"link" dc:"跳转链接"`
	Sort   int    `json:"sort" description:"sort" dc:"排序"`
}

type RotationRes struct {
	RotationId int64 `json:"rotation_id" example:"1" description:"rotation id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete/{Id}" tags:"轮播图" method:"delete" summary:"删除轮播图"`
	Id     int `json:"id" v:"required|min:1#轮播图id不能为空|轮播图id最小为1" dc:"轮播图id"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{Id}" tags:"轮播图" method:"put" summary:"更新轮播图"`
	Id     int    `json:"id" v:"required|min:1#轮播图id不能为空|轮播图id最小为1" dc:"轮播图id"`
	PicUrl string `json:"pic_url" v:"required:图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link" description:"link" dc:"跳转链接"`
	Sort   int    `json:"sort" description:"sort" dc:"排序"`
}

type RotationUpdateRes struct{}

type RotationListReq struct {
	g.Meta `path:"/backend/rotation/list" tags:"轮播图" method:"get" summary:"轮播图列表"`
	Sort   int `json:"sort" description:"sort" dc:"排序"`
	common.PaginationReq
}

type RotationListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
