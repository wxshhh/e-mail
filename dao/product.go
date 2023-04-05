package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

// NewProductDao 根据ctx初始化ProductDao
func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{DB: NewDBClient(ctx)}
}

// NewProductDaoByDB 利用已有DB创建ProductDao，实现DB复用，提升性能
func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{DB: db}
}

// CreateProduct 创建用户
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}

// CountProductByCondition 根据condition查询商品数目
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Offset(page.PageNum - 1).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset(page.PageNum - 1).Limit(page.PageSize).
		Find(&products).Error
	return
}

func (dao *ProductDao) GetProductByID(pid uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", pid).First(&product).Error
	return
}

func (dao *ProductDao) UpdateProductByID(id uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
}
