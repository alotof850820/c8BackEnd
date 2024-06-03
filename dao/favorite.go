package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

// 複製新DB
func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewFavoriteDaoByDb(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// 找全部Favorite
func (dao *FavoriteDao) GetFavorite(uId uint) (Favorite []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uId).Find(&Favorite).Error
	return
}

// 判斷是否已經收藏
func (dao *FavoriteDao) FavoriteExistOrNot(productId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", productId, uId).Count(&count).Error
	if err != nil {
		return false, err
	} else {
		if count > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

// 新增收藏
func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
	return
}

// 刪除收藏
func (dao *FavoriteDao) DeleteFavorite(uid, favoriteId uint) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("id = ? AND user_id = ?", favoriteId, uid).Delete(&model.Favorite{}).Error
	return
}
