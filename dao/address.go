package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

// 複製新DB
func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewAddressDaoByDb(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

// 找Address
func (dao *AddressDao) GetAddressByAId(aId uint) (Address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", aId).First(&Address).Error
	return
}

// 找全部的地址
func (dao *AddressDao) GetAddressesByUserId(uId uint) (Address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id = ?", uId).Find(&Address).Error
	return
}

// 新增地址
func (dao *AddressDao) CreateAddress(Address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Create(&Address).Error
	return
}

// 刪除地址
func (dao *AddressDao) DeleteAddressByAId(AddressId uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ? AND user_id = ? ", AddressId, uId).Delete(&model.Address{}).Error
	return
}

// 更新地址
func (dao *AddressDao) UpdateAddressByAId(AddressId uint, Address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ? ", AddressId).Updates(&Address).Error
	return
}
