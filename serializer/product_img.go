package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductID: item.ID,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		productImgs = append(productImgs, BuildProductImg(item))
	}
	return
}
