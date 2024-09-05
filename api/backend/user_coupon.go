package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" tags:"用户优惠券" method:"post" summary:"添加用户优惠券"`
	UserCouponCommonAddUpdate
}
type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id"      v:"required#用户id必填"   dc:"用户id"`
	CouponId uint  `json:"coupon_id"     v:"required#优惠券id必填" dc:"优惠券id"`
	Status   uint8 `json:"status"       dc:"状态"`
}
type UserCouponRes struct {
	UserCouponId uint `json:"coupon_id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update/{Id}" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint `json:"id"         v:"min:1#请选择需要修改的内容" dc:"用户优惠券Id"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct{}

type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user/coupon/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	//前后端分离不返回html
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
