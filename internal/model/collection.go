package model

import "github.com/gogf/gf/v2/frame/g"

type CollectionCreateInput struct {
	ObjectId int `json:"object_id"`
	UserId   int `json:"user_id"`
	Type     int `json:"type"`
}

type CollectionCreateOutput struct {
	Id int `json:"id"`
}

type CollectionCancelInput struct {
	Id       int `json:"id"`
	ObjectId int `json:"object_id"`
	Type     int `json:"type"`
}

type CollectionCancelOutput struct {
	Id int `json:"id"`
}

type CollectionListInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type CollectionItem struct {
	Id        int          `json:"id"`
	ObjectId  int          `json:"object_id"`
	UserId    int          `json:"user_id"`
	Type      int          `json:"type"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Goods     *GoodsItem   `json:"goods" orm:"with:id=object_id"`   // id为goods的关联字段
	Article   *ArticleItem `json:"article" orm:"with:id=object_id"` // id为article的关联字段
}

type CollectionListOutput struct {
	List  []*CollectionItem `json:"list"`
	Total int               `json:"total"`
}

type GoodsItem struct {
	g.Meta `orm:"table:goods_info"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"pic_url"`
	Price  int    `json:"price"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	PicUrl string `json:"pic_url"`
}
