package model

type RegisterInput struct {
	Name         string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Avatar       string `json:"avatar"       description:"头像"`
	Password     string `json:"password"     description:"密码" v:"password"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterOutput struct {
	Id uint
}

type LoginInput struct {
	Name     string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     description:"密码" v:"password"`
}
