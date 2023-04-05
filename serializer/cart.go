package serializer

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
)

type Cart struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	BossID        uint   `json:"boss_id"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossName      string `json:"boss_name"`
	CreateAt      int64  `json:"create_at"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		ID:            cart.ID,
		UserID:        cart.UserID,
		ProductID:     cart.ProductID,
		BossID:        boss.ID,
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: product.DiscountPrice,
		BossName:      boss.Username,
		CreateAt:      cart.CreatedAt.Unix(),
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) []Cart {
	var carts []Cart
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
		carts = append(carts, BuildCart(item, product, boss))
	}
	return carts
}
