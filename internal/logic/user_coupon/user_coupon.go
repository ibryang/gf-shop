package user_coupon

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sUserCoupon struct{}

func init() {
	// 注册路由
	service.RegisterUserCoupon(New())
}
func New() *sUserCoupon {
	return &sUserCoupon{}
}

// List 优惠券列表
func (s *sUserCoupon) List(ctx context.Context, in *model.UserCouponListInput) (out *model.UserCouponListItemOutput, err error) {
	var m = dao.UserCouponInfo.Ctx(ctx)

	out = &model.UserCouponListItemOutput{}
	listModel := m.Page(in.Page, in.PageSize)
	var list []*model.UserCouponListItem
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
func (s *sUserCoupon) Create(ctx context.Context, in *model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	id, err := dao.UserCouponInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.UserCouponCreateOutput{}, err
	}
	return model.UserCouponCreateOutput{Id: id}, nil
}

// Delete 优惠券删除
func (s *sUserCoupon) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.UserCouponInfo.Ctx(ctx).Delete(dao.UserCouponInfo.Columns().Id, id)
	return err
}

// Update 优惠券更新
func (s *sUserCoupon) Update(ctx context.Context, in *model.UserCouponUpdateInput) (err error) {
	_, err = dao.UserCouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.UserCouponInfo.Columns().Id).
		Where(dao.UserCouponInfo.Columns().Id, in.Id).
		Update()
	return err
}
