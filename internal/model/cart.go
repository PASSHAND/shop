package model

import (
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/model/entity"
)

type AddCartInput struct {
	UserId         uint
	GoodsOptionsId uint
	Count          uint
}

type AddCartOutput struct {
	Id uint `json:"id"`
}

type DeleteCartInput struct {
	Id uint
}

type DeleteCartOutput struct {
	Id uint `json:"id"`
}

type CartGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// CartGetListOutput 查询列表结果
type CartGetListOutput struct {
	List  []CartGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type CartGetListOutputItem struct {
	entity.CartInfo
}

// CartGetListOutput 查询列表结果
type CartDetailInput struct {
	Id uint
}

type CartDetailOutput struct {
	do.CartInfo
	Option do.GoodsOptionsInfo `orm:"with:id=goods_options_id"` //规格 sku
}
