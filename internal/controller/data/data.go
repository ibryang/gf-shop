package data

import (
	"context"
	"shop/api/backend"
	"shop/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) DataHead(ctx context.Context, req *backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	data, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	res = &backend.DataHeadRes{
		TodayDataCount: data.TodayDataCount,
		DAU:            data.DAU,
		ConversionRate: data.ConversionRate,
	}
	return
}
