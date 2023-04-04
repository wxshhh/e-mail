package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// NewCarouselDao 根据ctx初始化CarouseDao
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{DB: NewDBClient(ctx)}
}

// NewCarouselDaoByDB 利用已有DB创建CarouselDao，实现DB复用，提升性能
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{DB: db}
}

// ListCarousel 查询轮播图
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
