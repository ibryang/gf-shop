package file

import (
	"context"
	"fmt"
	"shop/api/common"
	"shop/internal/model"
	"shop/internal/service"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) List(ctx context.Context, req *common.FileListReq) (res *common.FileListRes, err error) {
	list, err := service.File().List(ctx, model.FileListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &common.FileListRes{
		Total: list.Total,
		List:  list.List,
	}, nil
}

func (c *Controller) Upload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error) {
	if req.File == nil {
		return nil, fmt.Errorf("上传文件不能为空")
	}
	out, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &common.FileUploadRes{Name: out.Name, Url: out.Url, Src: out.Src, Id: out.Id}, nil
}
