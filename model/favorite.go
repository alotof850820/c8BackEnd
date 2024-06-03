package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}

// `gorm:"size:256"`
//  `gorm:"check:age>30"`
//   `gorm:"size:256"`
//  `gorm:"scale:2;precision:7"` //小數點後2位 共7位
//  `gorm:"serializer:unixtime;type:time"`
//   `gorm:"serializer:json"`
