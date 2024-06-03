package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

// 複製新DB
func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewCategoryDaoByDb(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// 找全部Categorys
func (dao *CategoryDao) GetCategories() (categorys []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categorys).Error
	return
}
