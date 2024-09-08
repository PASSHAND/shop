package collection

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (out *model.AddCollectionOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddCollectionOutput{}, err
	}
	return &model.AddCollectionOutput{
		Id: gconv.Uint(id),
	}, err
}

// id为0按照对象和type删除
func (s *sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out *model.DeleteCollectionOutput, err error) {
	if in.Id != 0 {
		_, err := dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeleteCollectionOutput{Id: gconv.Uint(in.Id)}, err
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty().Where(in).Delete()
		if err != nil {
			return &model.DeleteCollectionOutput{}, err
		}
		return &model.DeleteCollectionOutput{
			Id: gconv.Uint(id),
		}, err
	}

}
