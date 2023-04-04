package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// NewNoticeDao 根据ctx初始化NoticeDao
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{DB: NewDBClient(ctx)}
}

// NewNoticeDaoByDB 利用已有DB创建UserDao，实现DB复用，提升性能
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{DB: db}
}

// GetNoticeByID 根据ID查找通知信息
func (dao *NoticeDao) GetNoticeByID(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return
}
