package praise

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) AddPraise(ctx context.Context, in model.AddPraiseInput) (out *model.AddPraiseOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddPraiseOutput{}, err
	}
	return &model.AddPraiseOutput{
		Id: gconv.Uint(id),
	}, err
}

// id为0按照对象和type删除
func (s *sPraise) DeletePraise(ctx context.Context, in model.DeletePraiseInput) (out *model.DeletePraiseOutput, err error) {
	if in.Id != 0 {
		_, err := dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeletePraiseOutput{Id: gconv.Uint(in.Id)}, err
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete()
		if err != nil {
			return &model.DeletePraiseOutput{}, err
		}
		return &model.DeletePraiseOutput{
			Id: gconv.Uint(id),
		}, err
	}

}

func (s *sPraise) GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	var (
		m = dao.PraiseInfo.Ctx(ctx)
	)
	out = &model.PraiseListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.PraiseListOutputItem{}, //数据为空时返回空数组
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	//只查询类型相同的
	if in.Type != 0 {
		listModel = listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}
	// Praise
	if in.Type == consts.PraiseTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
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

// 抽取获得点赞数量的方法
func PraiseCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     collectionType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// 抽取方法判断当前用户是否点赞
func CheckIsPraise(ctx context.Context, in model.CheckIsCollectionInput) (bool, error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
