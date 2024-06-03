package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

// 複製新DB
func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB(共享同一个连接池，从而提高性能和资源利用效率。)
func NewProductDaoByDb(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// 找全部商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}

// 根據條件找商品總數
func (dao *ProductDao) GetProductTotalByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// 根據條件找商品
func (dao *ProductDao) GetProductsByCondition(condition map[string]interface{}, basePage model.BasePage) (products []*model.Product, err error) {
	// Offset查询的偏移量 从当前页码-1开始查询。 Limit结果限制数量。
	err = dao.DB.Where(condition).Offset((basePage.PageNum - 1) * (basePage.PageSize)).Limit(basePage.PageSize).Find(&products).Error
	return
}

// 搜全部商品
func (dao *ProductDao) SearchProducts(info string, basePage model.BasePage) (products []*model.Product, total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where("title like ? or info like ?", "%"+info+"%", "%"+info+"%").
		Count(&total).Error
	if err != nil {
		return
	}
	// 根據info 模糊搜尋title or info
	err = dao.DB.Model(&model.Product{}).Where("title like ? or info like ?", "%"+info+"%", "%"+info+"%").
		Offset((basePage.PageNum - 1) * (basePage.PageSize)).
		Limit(basePage.PageSize).Find(&products).Error
	return
}

// 根據id找商品
func (dao *ProductDao) GetProductById(pid uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", pid).First(&product).Error
	return
}

// 更新商品
func (dao *ProductDao) UpdateProduct(pId uint, product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Where("id = ?", pId).Updates(product).Error
}
