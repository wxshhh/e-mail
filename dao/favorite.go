package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

// NewFavoriteDao 根据ctx初始化FavoriteDao
func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{DB: NewDBClient(ctx)}
}

// NewFavoriteDaoByDB 利用已有DB创建FavoriteDao，实现DB复用，提升性能
func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{DB: db}
}

// ListFavorite 查询分类
func (dao *FavoriteDao) ListFavorite(uid uint) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	return
}

func (dao *FavoriteDao) FavoriteExistOrNot(uid uint, pid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND product_id = ?", uid, pid).Count(&count).Error
	if err == nil && count > 0 {
		exist = true
	}
	return
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	return dao.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoriteDao) DeleteFavorite(uid, fid uint) error {
	return dao.Model(&model.Favorite{}).Where("user_id = ? AND id = ?", uid, fid).Delete(&model.Favorite{}).Error
}
