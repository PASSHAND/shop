package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

func (s *sUserCoupon) Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	lastInsertID, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCouponCreateOutput{UserCouponId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sUserCoupon) Delete(ctx context.Context, id uint) (err error) {
	// 删除内容
	_, err = dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
		dao.UserCouponInfo.Columns().Id: id,
	}).Delete() //加上.Unscoped()物理删除
	if err != nil {
		return err
	}
	return err

}

// Update 修改
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) error {
	_, err := dao.UserCouponInfo.
		Ctx(ctx).
		Data(in). //插入数据，in传进的数据会到data里面
		FieldsEx(dao.UserCouponInfo.Columns().Id).
		Where(dao.UserCouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var (
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.UserCouponInfo
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
	// UserCoupon
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
