package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

// 複製新DB
func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(productImg).Error
}

func (dao *ProductImgDao) GetProductImgs(id uint) (productImgs []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", id).Find(&productImgs).Error
	return
}
