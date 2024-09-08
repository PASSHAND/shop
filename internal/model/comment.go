package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop/internal/model/do"
)

type AddCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
	ParentId uint
	Content  string
}

type AddCommentOutput struct {
	Id uint `json:"id"`
}

type DeleteCommentInput struct {
	Id       uint
	UserId   uint
	Type     uint8
	ObjectId int
}

type DeleteCommentOutput struct {
	Id uint `json:"id"`
}

type CommentListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 类型
}

// CommentListOutput 查询列表结果
type CommentListOutput struct {
	List  []CommentListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type CommentListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	UserId    int         `json:"userId"  `
	ObjectId  int         `json:"objectId"`
	Type      int         `json:"type"    `
	ParentId  uint        `json:"goods"   dc:"父级评论id" `
	Content   string      `json:"article"      dc:"评论"`
	Goods     GoodsItem   `json:"parent_id" orm:"with:id=object_id"`
	Article   ArticleItem `json:"content"    orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CommentBase struct {
	do.CommentInfo
	User UserInfoBase `json:"user" orm:"with:id=user_id"`
}
