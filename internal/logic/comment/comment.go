package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) AddComment(ctx context.Context, in model.AddCommentInput) (out *model.AddCommentOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddCommentOutput{}, err
	}
	return &model.AddCommentOutput{
		Id: gconv.Uint(id),
	}, err
}

// id删除
func (s *sComment) DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out *model.DeleteCommentOutput, err error) {
	condition := g.Map{ //满足userid才可以删除
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	}
	_, err = dao.CommentInfo.Ctx(ctx).WherePri(condition).Delete()
	if err != nil {
		return nil, err
	}
	return &model.DeleteCommentOutput{Id: gconv.Uint(in.Id)}, err

}

func (s *sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	var (
		m = dao.CommentInfo.Ctx(ctx)
	)
	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CommentListOutputItem{}, //数据为空时返回空数组
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	//只查询类型相同的
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}
	// Comment
	if in.Type == consts.CommentTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
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

// 抽取获得评论数量的方法
func CommentCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     collectionType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// 抽取方法判断当前用户是否评论
func CheckIsComment(ctx context.Context, in model.CheckIsCollectionInput) (bool, error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
