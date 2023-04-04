package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserID"` // 一般不推荐使用外键
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"foreignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
