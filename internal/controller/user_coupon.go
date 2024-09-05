package controller

import (
	"context"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// UserCoupon 内容管理
var UserCoupon = cUserCoupon{}

type cUserCoupon struct{}

func (a *cUserCoupon) Create(ctx context.Context, req *backend.UserCouponReq) (res *backend.UserCouponRes, err error) {
	out, err := service.UserCoupon().Create(ctx, model.UserCouponCreateInput{
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.UserCouponRes{UserCouponId: out.UserCouponId}, nil
}

func (a *cUserCoupon) Delete(ctx context.Context, req *backend.UserCouponDeleteReq) (res *backend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, req.Id)
	return
}

func (a *cUserCoupon) Update(ctx context.Context, req *backend.UserCouponUpdateReq) (res *backend.UserCouponUpdateRes, err error) {
	err = service.UserCoupon().Update(ctx, model.UserCouponUpdateInput{
		Id: req.Id,
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	return
}

func (a *cUserCoupon) List(ctx context.Context, req *backend.UserCouponGetListCommonReq) (res *backend.UserCouponGetListCommonRes, err error) {
	getListRes, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{ //带着List出来
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.UserCouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
