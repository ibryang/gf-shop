package coupon

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sCoupon struct{}

func init() {
	// 注册路由
	service.RegisterCoupon(New())
}
func New() *sCoupon {
	return &sCoupon{}
}

// List 优惠券列表
func (s *sCoupon) List(ctx context.Context, in *model.CouponListInput) (out *model.CouponListItemOutput, err error) {
	var m = dao.CouponInfo.Ctx(ctx)

	out = &model.CouponListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)
	var list []*model.CouponListItem
	if err = listModel.Scan(&list); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return out, nil
	}

	out.Total, err = m.Count()
	if err != nil {
		return nil, err
	}
	out.List = list
	return out, nil
}

// Create 优惠券创建
func (s *sCoupon) Create(ctx context.Context, in *model.CouponCreateInput) (out model.CouponCreateOutput, err error) {
	id, err := dao.CouponInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.CouponCreateOutput{}, err
	}
	return model.CouponCreateOutput{CouponId: id}, nil
}

// Delete 优惠券删除
func (s *sCoupon) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.CouponInfo.Ctx(ctx).Delete(dao.CouponInfo.Columns().Id, id)
	return err
}

// Update 优惠券更新
func (s *sCoupon) Update(ctx context.Context, in *model.CouponUpdateInput) (err error) {
	_, err = dao.CouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CouponInfo.Columns().Id).
		Where(dao.CouponInfo.Columns().Id, in.Id).
		Update()
	return err
}
