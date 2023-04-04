package serializer

import (
	"gin_mall/model"
)

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"image_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	CreateAt      int64  `json:"create_at"`
}

func BuildFavorite(item *model.Favorite) Favorite {
	return Favorite{
		UserID:        item.UserID,
		ProductID:     item.ProductID,
		Name:          item.Product.Name,
		CategoryID:    item.Product.CategoryID,
		Title:         item.Product.Title,
		Info:          item.Product.Info,
		ImagePath:     item.Product.ImgPath,
		Price:         item.Product.Price,
		DiscountPrice: item.Product.DiscountPrice,
		BossID:        item.BossID,
		Num:           item.Product.Num,
		OnSale:        item.Product.OnSale,
		CreateAt:      item.CreatedAt.Unix(),
	}
}

func BuildFavorites(items []*model.Favorite) []Favorite {
	var favorites []Favorite
	for _, item := range items {
		favorites = append(favorites, BuildFavorite(item))
	}
	return favorites
}
