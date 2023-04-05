package serializer

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
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

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserID:        favorite.UserID,
		ProductID:     favorite.ProductID,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImagePath:     conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossID:        boss.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
		CreateAt:      favorite.CreatedAt.Unix(),
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) []Favorite {
	var favorites []Favorite
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductByID(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserByID(item.BossID)
		if err != nil {
			continue
		}
		favorites = append(favorites, BuildFavorite(item, product, boss))
	}
	return favorites
}
