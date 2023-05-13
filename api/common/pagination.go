package common

type PaginationReq struct {
	Page     int `json:"page" v:"required|min:1#页码不能为空|页码最小为1" dc:"页码" d:"1"`
	PageSize int `json:"page_size" v:"required|min:1#每页数量不能为空|每页数量最小为1" dc:"每页数量" d:"10"`
}
