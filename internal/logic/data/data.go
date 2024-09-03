package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sData struct {
}

func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
