package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/serializer"
	"strconv"
)

type PayService struct {
	OrderID   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductID uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossID    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       uint    `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"`
}

func (service *PayService) PayDown(ctx context.Context, uid uint) serializer.Response {
	utils.Encrypt.SetKey(service.Key)
	code := e.Success
	var err error
	orderDao := dao.NewOrderDao(ctx)
	// 开启事务
	tx := orderDao.Begin()
	order, err := orderDao.ShowOrder(service.OrderID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 修改对应商品数目
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(order.ProductID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	if product.Num < int(order.Num) {
		tx.Rollback()
		code = e.ErrorOutOfStock
		return serializer.ErrorByCode(code, err)
	}
	product.Num = product.Num - int(order.Num)
	err = productDao.UpdateProductByID(product.ID, product)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	price := order.Money
	num := order.Num
	price = price * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(order.BossID)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 对钱进行解密，减去订单金额，再进行保存
	userMoneyStr := utils.Encrypt.AesDecoding(user.Money)
	userMoney, _ := strconv.ParseFloat(userMoneyStr, 64)

	if userMoney < price {
		tx.Rollback()
		code = e.ErrorInsufficientFund
		return serializer.ErrorByCode(code, err)
	}

	// 更新用户金额
	user.Money = utils.Encrypt.AesEncoding(fmt.Sprintf("%f", userMoney-price))
	err = userDao.UpdateUserByID(user.ID, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	boss, err := userDao.GetUserByID(order.BossID)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 更新商家金额
	bossMoneyStr := utils.Encrypt.AesDecoding(boss.Money)
	bossMoney, _ := strconv.ParseFloat(bossMoneyStr, 64)
	boss.Money = utils.Encrypt.AesEncoding(fmt.Sprintf("%f", bossMoney+price))
	err = userDao.UpdateUserByID(boss.ID, boss)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 删除订单信息
	err = orderDao.DeleteOrder(order.ID)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	tx.Commit()

	return serializer.Success(nil)
}
