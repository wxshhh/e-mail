package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao 根据ctx初始化UserDao
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{DB: NewDBClient(ctx)}
}

// NewUserDaoByDB 利用已有DB创建UserDao，实现DB复用，提升性能
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{DB: db}
}

// ExistOrNotByUserName 根据username判断是否存在该名字的用户
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("username=?", username).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(user).Error
}

// GetUserByID 根据ID查找用户信息
func (dao *UserDao) GetUserByID(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", uid).First(&user).Error
	return
}

// UpdateUserByID 根据ID更新用户信息
func (dao *UserDao) UpdateUserByID(uid uint, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
	return
}
