package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
)

type FavoriteService struct {
	ProductID  uint `json:"product_id" form:"product_id"`
	BossID     uint `json:"boss_id" form:"boss_id"`
	FavoriteID uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, err := favoriteDao.FavoriteExistOrNot(uid, service.ProductID)
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.ErrorByCode(code, err)

	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	boss, err := userDao.GetUserByID(service.BossID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(service.ProductID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uid,
		Product:   *product,
		ProductID: service.ProductID,
		Boss:      *boss,
		BossID:    service.BossID,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}

func (service *FavoriteService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	var favorites []*model.Favorite
	favoriteDao := dao.NewFavoriteDao(ctx)
	favorites, err = favoriteDao.ListFavorite(uid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(len(favorites)))
}

func (service *FavoriteService) Delete(ctx context.Context, uid uint, id uint) serializer.Response {
	code := e.Success
	var err error
	favoriteDao := dao.NewFavoriteDao(ctx)
	err = favoriteDao.DeleteFavorite(uid, id)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}
