package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	uid, _ := c.Get("uid")
	var listOrder service.OrderService
	if err := c.ShouldBind(&listOrder); err == nil {
		res := listOrder.Create(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("CreateOrder api err:", err)
	}
}

func ListOrder(c *gin.Context) {
	var showOrder service.OrderService
	uid, _ := c.Get("uid")
	if err := c.ShouldBind(&showOrder); err == nil {
		res := showOrder.List(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListOrder api err:", err)
	}
}

func ShowOrder(c *gin.Context) {
	var createOrder service.OrderService
	aid, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&createOrder); err == nil {
		res := createOrder.Show(c.Request.Context(), uint(aid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ShowOrder api err:", err)
	}
}

func DeleteOrder(c *gin.Context) {
	var deleteOrder service.OrderService
	aid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&deleteOrder); err == nil {
		res := deleteOrder.Delete(c.Request.Context(), uint(aid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("DeleteOrder api err:", err)
	}
}
