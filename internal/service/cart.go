// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop/internal/model"
)

type (
	ICart interface {
		Add(ctx context.Context, in model.AddCartInput) (out model.AddCartOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, in model.DeleteCartInput) (out *model.DeleteCartOutput, err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CartGetListInput) (out *model.CartGetListOutput, err error)
		// 详情
		Detail(ctx context.Context, in model.CartDetailInput) (out model.CartDetailOutput, err error)
	}
)

var (
	localCart ICart
)

func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

func RegisterCart(i ICart) {
	localCart = i
}
