package category

import (
	"context"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sCategory) Delete(ctx context.Context, id uint) (err error) {
	// 删除内容
	_, err = dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Id: id,
	}).Delete() //加上.Unscoped()物理删除
	if err != nil {
		return err
	}
	return err

}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) error {
	_, err := dao.CategoryInfo.
		Ctx(ctx).
		Data(in). //插入数据，in传进的数据会到data里面
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Id)
	// 执行查询
	var list []*entity.CategoryInfo
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
	// Category
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
