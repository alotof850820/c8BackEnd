package model

import "gorm.io/gorm"

// 輪波圖案
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
