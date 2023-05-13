package model

type DataHeadOutput struct {
	TodayDataCount int `json:"today_data_count" desc:"今日订单总数"`
	DAU            int `json:"dau" desc:"今日日活"`
	ConversionRate int `json:"conversion_rate" desc:"转化率"`
}
