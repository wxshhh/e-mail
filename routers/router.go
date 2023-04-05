package routers

import (
	api "gin_mall/api/v1"
	"gin_mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct)
		v1.POST("products", api.SearchProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

		// 需要登录保护
		authed := v1.Group("/")
		{
			// 用户操作
			authed.Use(middleware.JWTAuth())
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendingEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 金额显示
			authed.GET("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)

			// 收藏夹操作
			authed.GET("favorites", api.ListFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			// 地址操作
			authed.GET("addresses", api.ListAddress)
			authed.GET("addresses/:id", api.ShowAddress)
			authed.POST("addresses", api.CreateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)

			// 购物车操作
			authed.GET("carts", api.ListCart)
			authed.GET("carts/:id", api.ShowCart)
			authed.POST("carts", api.CreateCart)
			authed.DELETE("carts/:id", api.DeleteCart)
			authed.PUT("carts/:id", api.UpdateCart)

			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.GET("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

			// 支付操作
			authed.POST("paydown", api.OrderPay)
		}
	}

	return r
}
