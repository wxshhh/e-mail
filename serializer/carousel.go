package serializer

import "gin_mall/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductID,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []model.Carousel) []Carousel {
	var carousels []Carousel
	for _, item := range items {
		carousels = append(carousels, BuildCarousel(&item))
	}
	return carousels
}
