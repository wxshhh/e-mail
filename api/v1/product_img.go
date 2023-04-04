package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProductImg(c *gin.Context) {
	var listProductImg service.ProductImgService
	if err := c.ShouldBind(&listProductImg); err == nil {
		res := listProductImg.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListCarousel api err:", err)
	}
}
