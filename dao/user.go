package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// 複製新DB
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根據userName是否存在該用戶
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	// 指定要操作的数据库表查询用戶是否存在将查询结果存储到 user 变量並获取是否错误
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

// 創建用戶
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

// 根據ID查詢用戶
func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", uid).First(&user).Error
	return user, err
}

// 依據用戶Id更新用戶
func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
}
