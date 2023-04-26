package model

// RotationCreateUpdateBase 轮播图创建或更新时模型
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
	Remark string
	Status int
}

// RotationCreateInput 轮播图创建时的输入参数
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 轮播图创建时的输出参数
type RotationCreateOutput struct {
	RequestId int64 `json:"request_id"`
}

// RotationUpdateInput 轮播图更新时的输入参数
type RotationUpdateInput struct {
	Id int
	RotationCreateUpdateBase
}
