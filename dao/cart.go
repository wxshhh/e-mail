package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

// NewCartDao 根据ctx初始化CartDao
func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{DB: NewDBClient(ctx)}
}

// NewCartDaoByDB 利用已有DB创建CartDao，实现DB复用，提升性能
func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{DB: db}
}

func (dao *CartDao) ListCart(uid uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uid).Find(&carts).Error
	return
}

func (dao *CartDao) ShowCart(cid uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", cid).First(&cart).Error
	return
}

func (dao *CartDao) CreateCart(cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&cart).Error
}

func (dao *CartDao) DeleteCart(cid uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", cid).Delete(&model.Cart{}).Error
	return
}

func (dao *CartDao) UpdateCart(cid uint, cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ?", cid).Updates(cart).Error
}
