package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
	CreateAt      int64  `json:"create_at"`
}

func BuildProduct(product *model.Product) Product {
	return Product{
		ID:            product.ID,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.View(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossID:        product.BossID,
		BossName:      product.BossName,
		BossAvatar:    conf.Host + conf.HttpPort + conf.AvatarPath + product.BossAvatar,
		CreateAt:      product.CreatedAt.Unix(),
	}
}

func BuildProducts(products []*model.Product) []Product {
	var productList []Product
	for _, product := range products {
		productList = append(productList, BuildProduct(product))
	}
	return productList
}
