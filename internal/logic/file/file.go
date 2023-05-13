package file

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
)

type sFile struct {
}

func New() *sFile {
	return &sFile{}
}

func init() {
	// 注册路由
	service.RegisterFile(New())
}

// Upload
//
//	1.定义上传位置
//	2.校验上传文件是否正确
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out model.FiLeUploadOutput, err error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	dateDirname := gtime.Now().Format("Ymd")
	filename, err := in.File.Save(gfile.Join(uploadPath, dateDirname), in.RandomName)
	if err != nil {
		return model.FiLeUploadOutput{}, err
	}
	url := gfile.Join(uploadPath, dateDirname, filename)
	// 入库
	id, err := dao.FileInfo.Ctx(ctx).OmitEmpty().InsertAndGetId(&entity.FileInfo{
		Name:   filename,
		Src:    in.File.Filename,
		Url:    url,
		UserId: 9,
	})
	if err != nil {
		return model.FiLeUploadOutput{}, err
	}
	return model.FiLeUploadOutput{Name: filename, Url: url, Id: int(id), Src: in.File.Filename}, nil
}

func (s *sFile) List(ctx context.Context, in model.FileListInput) (out *model.FileListOutput, err error) {
	var files []*entity.FileInfo
	m := dao.FileInfo.Ctx(ctx)
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	err = m.Page(in.Page, in.PageSize).Scan(&files)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	out = &model.FileListOutput{
		List:  files,
		Total: count,
	}
	return out, nil
}
