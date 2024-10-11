package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Cart 内容管理
var Cart = cCart{}

type cCart struct{}

func (a *cCart) Add(ctx context.Context, req *frontend.AddCartReq) (res *frontend.AddCartRes, err error) {
	data := model.AddCartInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Cart().Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCartRes{Id: out.Id}, nil
}

func (a *cCart) Delete(ctx context.Context, req *frontend.DeleteCartReq) (res *frontend.DeleteCartRes, err error) {
	data := model.DeleteCartInput{Id: req.Id}
	out, err := service.Cart().Delete(ctx, data)
	return &frontend.DeleteCartRes{Id: out.Id}, nil
}

func (a *cCart) List(ctx context.Context, req *frontend.ListCartReq) (res *frontend.ListCartRes, err error) {
	getListRes, err := service.Cart().GetList(ctx, model.CartGetListInput{ //带着List出来
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &frontend.ListCartRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (a *cCart) Detail(ctx context.Context, req *frontend.CartDetailReq) (res *frontend.CartDetailRes, err error) {
	detail, err := service.Cart().Detail(ctx, model.CartDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &frontend.CartDetailRes{}
	err = gconv.Struct(detail, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
