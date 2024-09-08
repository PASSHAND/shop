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

func (s *sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	var (
		m = dao.CollectionInfo.Ctx(ctx)
	)
	out = &model.CollectionListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CollectionListOutputItem{}, //数据为空时返回空数组
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	//只查询类型相同的
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}
	// Collection
	if in.Type == consts.CollectionTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CollectionTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}

	return
}
