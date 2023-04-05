package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service AddressService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	var addresses []*model.Address
	addressDao := dao.NewAddressDao(ctx)
	addresses, err = addressDao.ListAddress(uid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(len(addresses)))
}

func (service AddressService) Show(ctx context.Context, aid uint) serializer.Response {
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.ShowAddress(aid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildAddress(address))
}

func (service AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(address)
}

func (service AddressService) Delete(ctx context.Context, aid uint) serializer.Response {
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)
	err = addressDao.DeleteAddress(aid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}

func (service AddressService) Update(ctx context.Context, uid uint, aid uint) serializer.Response {
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err = addressDao.UpdateAddress(aid, address)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildAddress(address))
}
