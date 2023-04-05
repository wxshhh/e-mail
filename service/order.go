package service

import (
	"context"
	"fmt"
	"gin_mall/cache"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	UserID    uint `json:"user_id" form:"user_id"`
	ProductID uint `json:"product_id" form:"product_id"`
	BossID    uint `json:"boss_id" form:"boss_id"`
	AddressID uint `json:"address_id" form:"address_id"`
	Num       uint `json:"num" form:"num"`
	OrderNum  uint `json:"order_num" form:"order_num"`
	Type      uint `json:"type" form:"type"`
	Money     int  `json:"money" form:"money"`
	model.BasePage
}

func (service OrderService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error

	order := &model.Order{
		UserID:    uid,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       service.Num,
		Type:      1,
		Money:     float64(service.Money),
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.ShowAddress(service.AddressID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	order.AddressID = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(service.UserID))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.CreateOrder(order)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 订单号存入Redis中，设置过期时间
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: orderNum,
	}
	cache.RedisClient.ZAdd(cache.OrderTimeKey, data)

	return serializer.Success(order)
}

func (service OrderService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var err error
	var orders []*model.Order
	var total int64
	if service.PageSize == 0 {
		service.PageSize = 5
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	condition["user_id"] = uid
	condition["type"] = service.Type

	orders, total, err = orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))
}

func (service OrderService) Show(ctx context.Context, oid uint) serializer.Response {
	code := e.Success
	var err error

	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.ShowOrder(oid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(order.ProductID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.ShowAddress(order.AddressID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	return serializer.Success(serializer.BuildOrder(order, product, address))
}

func (service OrderService) Delete(ctx context.Context, oid uint) serializer.Response {
	code := e.Success
	var err error
	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.DeleteOrder(oid)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}
