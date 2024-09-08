package model

type AddCollectionInput struct {
	UserId   uint  `json:"user_id"     description:"用户id"`
	ObjectId uint  `json:"object_id"   description:"对象id" v:"required#对象id必填"`
	Type     uint8 `json:"type"       description:"收藏类型：1商品 2文章" v:"in：1,2"`
}

type AddCollectionOutput struct {
	Id uint `json:"id"`
}

type DeleteCollectionInput struct {
	Id       uint  `json:"id"`
	UserId   uint  `json:"user_id"   `
	Type     uint8 `json:"type"`
	ObjectId int   `json:"object_id"`
}

type DeleteCollectionOutput struct {
	Id uint `json:"id"`
}
