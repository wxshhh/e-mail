package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

// NewOrderDao 根据ctx初始化OrderDao
func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{DB: NewDBClient(ctx)}
}

// NewOrderDaoByDB 利用已有DB创建OrderDao，实现DB复用，提升性能
func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{DB: db}
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Offset(page.PageNum - 1).Limit(page.PageSize).Find(&orders).Error
	return orders, total, err
}

func (dao *OrderDao) ShowOrder(oid uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ?", oid).First(&order).Error
	return
}

func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&order).Error
}

func (dao *OrderDao) DeleteOrder(oid uint) (err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ?", oid).Delete(&model.Order{}).Error
	return
}

func (dao *OrderDao) UpdateOrder(oid uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id = ?", oid).Updates(order).Error
}
