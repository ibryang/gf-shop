package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

// 数据大屏
type sData struct{}

func New() *sData {
	return &sData{}
}

func init() {
	service.RegisterData(New())
}

func (s *sData) DataHead(ctx context.Context) (output *model.DataHeadOutput, err error) {
	output = &model.DataHeadOutput{
		TodayDataCount: getTodayDataCount(ctx),
		DAU:            grand.N(0, 999999),
		ConversionRate: grand.N(0, 50),
	}
	return
}

// 查询今天订单总数
func getTodayDataCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return
}
