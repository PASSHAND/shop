package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"文章" method:"post" summary:"添加文章"`
	ArticleCommonAddUpdate
}
type ArticleCommonAddUpdate struct {
	//UserId  int    `json:"userId"     dc:"用户id" v:"required#用户id必填"`
	Title   string `json:"title"      description:"标题" v:"required#标题必填"`
	Desc    string `json:"desc"       description:"摘要"`
	PicUrl  string `json:"picUrl"     description:"封面图"`
	IsAdmin int    `description:"1后台管理员发布 2前台用户发布" d:"1"` //默认值1，不需要传值
	Praise  int    `json:"praise"     description:"点赞数"`
	Detail  string `json:"detail"    orm:"detail"     description:"文章详情" v:"required#文章详情必填"`
}
type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"文章" summary:"删除文章接口"`
	Id     uint `v:"min:1#请选择需要删除的文章" dc:"文章id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update/{Id}" method:"post" tags:"文章" summary:"修改文章接口"`
	Id     uint `json:"id"         v:"min:1#请选择需要修改的内容" dc:"文章Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}

type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章" summary:"文章列表接口"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	//前后端分离不返回html
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
