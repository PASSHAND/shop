package controller

import (
	"context"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Category 内容管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryRes{CategoryId: out.CategoryId}, nil
}

func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: req.Id,
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
		},
	})
	return
}

func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{ //带着List出来
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CategoryGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
