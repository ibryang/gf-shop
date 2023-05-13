package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" tags:"文件管理" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"urL" dc:"访问URL,可能只是URI"`
	Src  string `json:"src"`
	Id   int    `json:"id"`
}

type FileListReq struct {
	g.Meta `path:"/file/list" method:"get" tags:"文件管理" summary:"文件列表"`
	PaginationReq
}

type FileListRes struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
