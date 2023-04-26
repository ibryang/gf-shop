// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IPosition interface {
		List(ctx context.Context, in *model.PositionListInput) (out *model.PositionListItemOutput, err error)
		Create(ctx context.Context, in *model.PositionCreateInput) (out model.PositionCreateOutput, err error)
		Delete(ctx context.Context, id int) (err error)
		Update(ctx context.Context, in *model.PositionUpdateInput) (err error)
	}
)

var (
	localPosition IPosition
)

func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}

func RegisterPosition(i IPosition) {
	localPosition = i
}
