package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

// 複製新DB
func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewCartDaoByDb(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

// 找Cart
func (dao *CartDao) GetCartByCId(cId uint) (Cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", cId).First(&Cart).Error
	return
}

// 找全部的購物車
func (dao *CartDao) GetCartesByUserId(uId uint) (Cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&Cart).Error
	return
}

// 新增購物車
func (dao *CartDao) CreateCart(Cart *model.Cart) (err error) {
	err = dao.DB.Model(&model.Cart{}).Create(&Cart).Error
	return
}

// 刪除購物車
func (dao *CartDao) DeleteCartByCId(CartId uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ? ", CartId, uId).Delete(&model.Cart{}).Error
	return
}

// 更新購物車
func (dao *CartDao) UpdateCartByCId(CartId uint, Cart *model.Cart) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ? ", CartId).Updates(&Cart).Error
	return
}

// 更新購物車數量
func (dao *CartDao) UpdateCartNumByCId(CartId uint, num int) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ? ", CartId).Update("num", num).Error
	return
}
