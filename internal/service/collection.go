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
	ICollection interface {
		AddCollection(ctx context.Context, in model.AddCollectionInput) (out *model.AddCollectionOutput, err error)
		// id为0按照对象和type删除
		DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out *model.DeleteCollectionOutput, err error)
	}
)

var (
	localCollection ICollection
)

func Collection() ICollection {
	if localCollection == nil {
		panic("implement not found for interface ICollection, forgot register?")
	}
	return localCollection
}

func RegisterCollection(i ICollection) {
	localCollection = i
}
