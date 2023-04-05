package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListAddress(c *gin.Context) {
	uid, _ := c.Get("uid")
	var listAddress service.AddressService
	if err := c.ShouldBind(&listAddress); err == nil {
		res := listAddress.List(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListAddress api err:", err)
	}
}

func ShowAddress(c *gin.Context) {
	var showAddress service.AddressService
	aid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&showAddress); err == nil {
		res := showAddress.Show(c.Request.Context(), uint(aid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListAddress api err:", err)
	}
}

func CreateAddress(c *gin.Context) {
	uid, _ := c.Get("uid")
	var createAddress service.AddressService
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Create(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListAddress api err:", err)
	}
}

func DeleteAddress(c *gin.Context) {
	var deleteAddress service.AddressService
	aid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&deleteAddress); err == nil {
		res := deleteAddress.Delete(c.Request.Context(), uint(aid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListAddress api err:", err)
	}
}

func UpdateAddress(c *gin.Context) {
	uid, _ := c.Get("uid")
	var updateAddress service.AddressService
	aid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&updateAddress); err == nil {
		res := updateAddress.Update(c.Request.Context(), uid.(uint), uint(aid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListAddress api err:", err)
	}
}
