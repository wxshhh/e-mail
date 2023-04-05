package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListCart(c *gin.Context) {
	uid, _ := c.Get("uid")
	var listCart service.CartService
	if err := c.ShouldBind(&listCart); err == nil {
		res := listCart.List(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCart api err:", err)
	}
}

func ShowCart(c *gin.Context) {
	var showCart service.CartService
	cid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&showCart); err == nil {
		res := showCart.Show(c.Request.Context(), uint(cid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCart api err:", err)
	}
}
func CreateCart(c *gin.Context) {
	uid, _ := c.Get("uid")
	var createCart service.CartService
	if err := c.ShouldBind(&createCart); err == nil {
		res := createCart.Create(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCart api err:", err)
	}
}
func DeleteCart(c *gin.Context) {
	var deleteCart service.CartService
	cid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&deleteCart); err == nil {
		res := deleteCart.Delete(c.Request.Context(), uint(cid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCart api err:", err)
	}
}
func UpdateCart(c *gin.Context) {
	uid, _ := c.Get("uid")
	var updateCart service.CartService
	cid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&updateCart); err == nil {
		res := updateCart.Update(c.Request.Context(), uid.(uint), uint(cid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCart api err:", err)
	}
}
