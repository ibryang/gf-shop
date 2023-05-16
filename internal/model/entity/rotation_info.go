// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure for table rotation_info.
type RotationInfo struct {
	Id        int         `json:"id"         description:""`
	PicUrl    string      `json:"pic_url"    description:"轮播图片"`
	Link      string      `json:"link"       description:"跳转链接"`
	Sort      int         `json:"sort"       description:"排序字段"`
	CreatedAt *gtime.Time `json:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deleted_at" description:""`
}
