package dao

import (
	"gin_mall_tmp/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// 複製新DB
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{
		NewDBClient(&ctx),
	}
}

// 複製已有的DB
func NewNoticeDaoByDb(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// 根據ID查詢notice
func (dao *NoticeDao) GetNoticeById(uid uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id = ?", uid).First(&notice).Error
	return notice, err
}
