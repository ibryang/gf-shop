package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadInput struct {
	File       *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
	Name       string
	RandomName bool
}

type FiLeUploadOutput struct {
	Id   int    `json:"id"`
	Src  string `json:"src"`
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"访问URL,可能只是URI"`
}

type FileUploadsInput struct {
	Files *ghttp.UploadFiles `json:"files" type:"file" dc:"选择上传文件"`
}

type FiLeUploadsOutput struct {
	List interface{} `json:"list"`
}

type FileListInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type FileListOutput struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
