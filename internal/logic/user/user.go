package user

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//加密盐逻辑，生成一个随机字符串，长度为n
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt) //存储一同加密的密码
	in.UserSalt = UserSalt                                       //存储随机数

	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}
