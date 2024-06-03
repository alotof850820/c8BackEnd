package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CarouselsDao struct {
	*gorm.DB
}

// 複製新DB
func NewCarouselsDao(ctx context.Context) *CarouselsDao {
	return &CarouselsDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewCarouselsDaoByDb(db *gorm.DB) *CarouselsDao {
	return &CarouselsDao{db}
}

// 找全部carousels
func (dao *CarouselsDao) GetListCarousels() (carousels []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return carousels, err
}
