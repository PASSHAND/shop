package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryCreateUpdateBase 创建/修改内容基类
type CategoryCreateUpdateBase struct {
	ParentId int
	Name     string
	PicUrl   string
	Level    uint8
	Sort     uint8
}

// CategoryCreateInput 创建内容
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
}

// CategoryCreateOutput 创建内容返回结果
type CategoryCreateOutput struct {
	CategoryId int `json:"category_id"`
}

// CategoryUpdateInput 修改内容
type CategoryUpdateInput struct {
	CategoryCreateUpdateBase
	Id uint
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type CategoryGetListOutputItem struct {
	Id        uint        `json:"id"`
	ParentId  int         `json:"parent_id"` // 自增ID
	Name      string      `json:"name"`
	PicUrl    string      `json:"PicUrl"`
	Level     uint8       `json:"level"`
	Sort      uint8       `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CategorySearchOutputItem struct {
	CategoryGetListOutputItem
}
