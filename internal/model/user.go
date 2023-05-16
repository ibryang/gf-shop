package model

type UserRegisterInput struct {
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Password     string `json:"password"`
	UserSalt     string `json:"user_salt"`
	Sex          int    `json:"sex"`
	Status       int    `json:"status"`
	Sign         string `json:"sign"`
	SecretAnswer string `json:"secret_answer"`
}

type UserRegisterOutput struct {
	Id int64 `json:"id"`
}
