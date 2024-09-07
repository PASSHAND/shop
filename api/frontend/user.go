package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta       `path:"/register" method:"post" tags:"前台注册" summary:"前台用户注册接口"`
	Name         string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Avatar       string `json:"avatar"       description:"头像"`
	Password     string `json:"password"     description:"密码" v:"password"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterRes struct {
	Id uint `json:"id"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"前台登录" summary:"前台用户登录接口"`
	Name     string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     description:"密码" v:"password"`
}

// for gtoken
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Name     string `json:"name"`   //姓名
	Avatar   string `json:"avatar"` //头像
	Sex      uint8  `json:"sex"`
	Sign     string `json:"sign"`
	Status   uint8  `json:"status"`
}
