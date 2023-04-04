package dao

import (
	"fmt"
	"gin_mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Product{},
			&model.Category{},
			&model.Favorite{},
			&model.Order{},
			&model.Notice{},
			&model.Carousel{},
			&model.ProductImg{},
			&model.Admin{},
			&model.Cart{},
			&model.Address{},
		)
	if err != nil {
		fmt.Println("err: ", err)
	}
	return
}
