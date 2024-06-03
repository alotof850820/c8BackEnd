package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

// 複製新DB
func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewOrderDaoByDb(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

// 找Order
func (dao *OrderDao) GetOrderByOId(oId uint, uId uint) (Order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oId, uId).First(&Order).Error
	return
}

// 找全部的訂單
func (dao *OrderDao) GetOrderesByCondition(condition map[string]interface{}, page model.BasePage) (Order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&Order).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	return
}

// 新增訂單
func (dao *OrderDao) CreateOrder(Order *model.Order) (err error) {
	err = dao.DB.Model(&model.Order{}).Create(&Order).Error
	return
}

// 刪除訂單
func (dao *OrderDao) DeleteOrderByOId(OrderId uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ? ", OrderId, uId).Delete(&model.Order{}).Error
	return
}
