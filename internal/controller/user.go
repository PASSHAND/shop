package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// User 内容管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{Id: out.Id}, nil
}
