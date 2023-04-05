package serializer

import "gin_mall/model"

type Address struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) []Address {
	var addresses []Address
	for _, item := range items {
		addresses = append(addresses, BuildAddress(item))
	}
	return addresses
}
