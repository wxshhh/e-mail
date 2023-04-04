package serializer

import (
	"gin_mall/model"
	"gin_mall/pkg/utils"
)

type Money struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) *Money {
	utils.Encrypt.SetKey(key)
	return &Money{
		UserID:    user.ID,
		UserName:  user.Username,
		UserMoney: utils.Encrypt.AesDecoding(user.Money),
	}
}
