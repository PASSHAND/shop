package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionCreateUpdateBase struct {
	Name string `json:"name"       v:"required#名称不能为空" dc:"权限名称"`
	Path string `json:"path"       dc:"权限路径"`
}
type PermissionReq struct {
	g.Meta `path:"/backend/permission/add" method:"post" desc:"添加权限" tags:"permission"`
	PermissionCreateUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

// 更新
type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update" method:"post" desc:"修改权限" tags:"permission"`
	Id     uint `json:"id"         v:"required#id必填" desc:"id"`
	PermissionCreateUpdateBase
}

type PermissionUpdateRes struct {
	Id uint `json:"id"`
}

// 删除
type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete" method:"delete" tags:"权限" summary:"删除权限接口"`
	Id     uint `v:"min:1#请选择需要删除的权限" dc:"权限id"`
}
type PermissionDeleteRes struct{}

// 列表
type PermissionGetListCommonReq struct {
	g.Meta `path:"/backend/permission/list" method:"get" tags:"权限" summary:"权限列表"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	//前后端分离不返回html
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
