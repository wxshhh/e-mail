package serializer

import "gin_mall/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) []Category {
	var categories []Category
	for _, item := range items {
		categories = append(categories, BuildCategory(item))
	}
	return categories
}
