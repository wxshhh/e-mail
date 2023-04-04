package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	uid, _ := c.Get("uid")
	var createProduct service.ProductService
	if err := c.ShouldBind(&createProduct); err == nil {
		res := createProduct.Create(c.Request.Context(), uid.(uint), files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("CreateProduct api err:", err)
	}
}

// ListProduct 分页查询商品
func ListProduct(c *gin.Context) {
	var listProduct service.ProductService
	if err := c.ShouldBind(&listProduct); err == nil {
		res := listProduct.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListProduct api err:", err)
	}
}

// SearchProduct 根据关键词搜索商品
func SearchProduct(c *gin.Context) {
	var searchProduct service.ProductService
	if err := c.ShouldBind(&searchProduct); err == nil {
		res := searchProduct.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("SearchProduct api err:", err)
	}
}

// ShowProduct 根据ID搜索商品
func ShowProduct(c *gin.Context) {
	var showProduct service.ProductService
	if err := c.ShouldBind(&showProduct); err == nil {
		res := showProduct.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ShowProduct api err:", err)
	}
}
