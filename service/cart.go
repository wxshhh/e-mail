package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
)

type CartService struct {
	ID        uint `json:"id" form:"id"`
	ProductID uint `json:"product_id" form:"product_id"`
	BossID    uint `json:"boss_id" form:"boss_id"`
	Num       uint `json:"num" form:"num"`
}

func (service CartService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	var carts []*model.Cart
	cartDao := dao.NewCartDao(ctx)
	carts, err = cartDao.ListCart(uid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.BuildListResponse(serializer.BuildCarts(ctx, carts), uint(len(carts)))
}

func (service CartService) Show(ctx context.Context, cid uint) serializer.Response {
	code := e.Success
	var err error

	cartDao := dao.NewCartDao(ctx)
	cart, err := cartDao.ShowCart(cid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(cart.ProductID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserByID(cart.BossID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	return serializer.Success(serializer.BuildCart(cart, product, boss))
}

func (service CartService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(service.ProductID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	cartDao := dao.NewCartDao(ctx)
	cart := &model.Cart{
		UserID:    uid,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       service.Num,
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserByID(service.BossID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildCart(cart, product, boss))
}

func (service CartService) Delete(ctx context.Context, cid uint) serializer.Response {
	code := e.Success
	var err error
	cartDao := dao.NewCartDao(ctx)
	err = cartDao.DeleteCart(cid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}

func (service CartService) Update(ctx context.Context, uid uint, cid uint) serializer.Response {
	code := e.Success
	var err error
	cartDao := dao.NewCartDao(ctx)
	cart := &model.Cart{
		UserID:    uid,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       service.Num,
	}
	err = cartDao.UpdateCart(cid, cart)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}
