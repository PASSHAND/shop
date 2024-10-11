package cart

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sCart struct{}

func init() {
	service.RegisterCart(New())
}

func New() *sCart {
	return &sCart{}
}

func (s *sCart) Add(ctx context.Context, in model.AddCartInput) (out model.AddCartOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.CartInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddCartOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sCart) Delete(ctx context.Context, in model.DeleteCartInput) (out *model.DeleteCartOutput, err error) {
	// 删除内容
	_, err = dao.CartInfo.Ctx(ctx).Where(g.Map{
		dao.CartInfo.Columns().Id: in.Id,
	}).Delete() //加上.Unscoped()物理删除
	if err != nil {
		return nil, err
	}
	return &model.DeleteCartOutput{Id: gconv.Uint(in.Id)}, err

}

// GetList 查询分类列表
func (s *sCart) GetList(ctx context.Context, in model.CartGetListInput) (out *model.CartGetListOutput, err error) {
	var (
		m = dao.CartInfo.Ctx(ctx)
	)
	out = &model.CartGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.CartInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Cart
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// 详情
func (*sCart) Detail(ctx context.Context, in model.CartDetailInput) (out model.CartDetailOutput, err error) {
	err = dao.CartInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return out, err
	}
	return
}
