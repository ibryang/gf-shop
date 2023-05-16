package model

type CollectionCreateInput struct {
	ObjectId int `json:"object_id"`
	UserId   int `json:"user_id"`
	Type     int `json:"type"`
}

type CollectionCreateOutput struct {
	Id int `json:"id"`
}

type CollectionCancelInput struct {
	Id       int `json:"id"`
	ObjectId int `json:"object_id"`
	Type     int `json:"type"`
}

type CollectionCancelOutput struct {
	Id int `json:"id"`
}
