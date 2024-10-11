package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/entity"
)

type AddCartReq struct {
	g.Meta         `path:"/add/cart" method:"post" tags:"购物车" summary:"添加购物车"`
	GoodsOptionsId uint `json:"goods_options_id"   dc:"对象id" v:"required#商品规格id必填"`
	Count          uint `json:"count" dc:"数量" v:"required#商品数量必填"`
}

type AddCartRes struct {
	Id uint `json:"id"`
}

// 只能按照id删除
type DeleteCartReq struct {
	g.Meta `path:"/delete/cart" method:"post" tags:"购物车" summary:"删除购物车"`
	Id     uint `json:"id" v:"min:1#请选择需要删除的商品" dc:"购物车商品id"`
}

type DeleteCartRes struct {
	Id uint `json:"id"`
}

type ListCartReq struct {
	g.Meta `path:"/cart/list" method:"post" tags:"购物车" summary:"购物车列表"`
	CommonPaginationReq
}

type ListCartRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type CartDetailReq struct {
	g.Meta `path:"/cart/detail" method:"post" tags:"购物车" summary:"购物车详情"`
	Id     uint `json:"id"`
}

type CartDetailRes struct {
	entity.CartInfo             //会变成驼峰命名
	Option          interface{} `json:"option"` //规格 sku
}
