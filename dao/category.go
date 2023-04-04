package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

// NewCategoryDao 根据ctx初始化CategoryDao
func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{DB: NewDBClient(ctx)}
}

// NewCategoryDaoByDB 利用已有DB创建CategoryDao，实现DB复用，提升性能
func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{DB: db}
}

// ListCategory 查询分类
func (dao *CategoryDao) ListCategory() (categories []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categories).Error
	return
}
