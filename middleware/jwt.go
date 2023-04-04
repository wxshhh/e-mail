package middleware

import (
	"fmt"
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var userID uint
		code = e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthToken
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil || claims == nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			} else {
				userID = claims.ID
			}
		}
		if code != e.Success {
			fmt.Println("---------------", code)
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("uid", userID)
		c.Next()
	}
}
