package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

// NewProductImgDao 根据ctx初始化ProductImgDao
func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{DB: NewDBClient(ctx)}
}

// NewProductImgDaoByDB 利用已有DB创建ProductImgDao，实现DB复用，提升性能
func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{DB: db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(productImg).Error
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", id).Find(&productImg).Error
	return
}
