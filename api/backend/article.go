package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/api/common"
)

type ArticleAddUpdateBaseReq struct {
	UserId  int    `json:"user_id" v:"required#用户id不能为空" dc:"用户id"`
	Title   string `json:"title" v:"required#文章名称不能为空" dc:"文章名称"`
	Desc    string `json:"desc" v:"" dc:"文章概要"`
	PicUrl  string `json:"pic_url" dc:"图片"`
	IsAdmin int    `json:"is_admin" v:"" dc:"是否是管理员发布" d:"1"`
	Detail  string `json:"detail" v:"required#文章详情" dc:"文章详情"`
	Praise  int    `json:"praise" v:"" dc:"点赞数"`
}

type ArticleAddReq struct {
	g.Meta `path:"/backend/article/add" tags:"文章" method:"post" summary:"添加文章"`
	ArticleAddUpdateBaseReq
}

type ArticleRes struct {
	Id int64 `json:"id" example:"1" dc:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/backend/article/delete/{Id}" tags:"文章" method:"delete" summary:"删除文章"`
	Id     int `json:"id" v:"required|min:1#文章id不能为空|文章id最小为1" dc:"文章id"`
}

type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/backend/article/update/{Id}" tags:"文章" method:"put" summary:"更新文章"`
	Id     int `json:"id" v:"required|min:1#文章id不能为空|文章id最小为1" dc:"文章id"`
	ArticleAddUpdateBaseReq
}

type ArticleUpdateRes struct{}

type ArticleListReq struct {
	g.Meta `path:"/backend/article/list" tags:"文章" method:"get" summary:"文章列表"`
	Sort   int `json:"sort" dc:"sort" dc:"排序"`
	common.PaginationReq
}

type ArticleListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
