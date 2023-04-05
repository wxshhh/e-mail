package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderPay(c *gin.Context) {
	var orderPay service.PayService
	uid, _ := c.Get("uid")
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("OrderPay api err:", err)
	}
}
