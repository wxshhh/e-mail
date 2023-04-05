package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

// NewAddressDao 根据ctx初始化AddressDao
func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{DB: NewDBClient(ctx)}
}

// NewAddressDaoByDB 利用已有DB创建AddressDao，实现DB复用，提升性能
func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{DB: db}
}

// ListAddress 查询分类
func (dao *AddressDao) ListAddress(uid uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id = ?", uid).Find(&addresses).Error
	return
}

func (dao *AddressDao) ShowAddress(aid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", aid).First(&address).Error
	return
}

func (dao *AddressDao) CreateAddress(address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

func (dao *AddressDao) DeleteAddress(aid uint) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", aid).Delete(&model.Address{}).Error
	return
}

func (dao *AddressDao) UpdateAddress(aid uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id = ?", aid).Updates(address).Error
}
