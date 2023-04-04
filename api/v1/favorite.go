package v1

import (
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateFavorite 创建收藏夹
func CreateFavorite(c *gin.Context) {
	uid, _ := c.Get("uid")
	var createFavorite service.FavoriteService
	if err := c.ShouldBind(&createFavorite); err == nil {
		res := createFavorite.Create(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("CreateFavorite api err:", err)
	}
}

// ListFavorite 分页查询收藏夹
func ListFavorite(c *gin.Context) {
	uid, _ := c.Get("uid")
	var listFavorite service.FavoriteService
	if err := c.ShouldBind(&listFavorite); err == nil {
		res := listFavorite.List(c.Request.Context(), uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("ListFavorite api err:", err)
	}
}

// DeleteFavorite 根据ID搜索收藏夹
func DeleteFavorite(c *gin.Context) {
	uid, _ := c.Get("uid")
	var deleteFavorite service.FavoriteService
	var err error
	fid, err := strconv.Atoi(c.Param("id"))
	if err = c.ShouldBind(&deleteFavorite); err == nil {
		res := deleteFavorite.Delete(c.Request.Context(), uid.(uint), uint(fid))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln("DeleteFavorite api err:", err)
	}
}
