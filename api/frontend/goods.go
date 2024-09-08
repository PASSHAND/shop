package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/entity"
)

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品" summary:"商品列表接口"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	//前后端分离不返回html
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"post" tags:"前台商品" summary:"商品详情"`
	Id     uint `json:"id"`
}

type GoodsDetailRes struct {
	entity.GoodsInfo             //会变成驼峰命名
	Option           interface{} `json:"option"` //规格 sku
	Comment          interface{} `json:"comment"`
	IsCollect        bool        `json:"is_collect"`
}
