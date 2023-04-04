package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`
	AddressID uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	OrderNum  uint64
	Type      uint // 1：未支付 2：已支付
	Money     float64
}
