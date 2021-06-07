package middleware

import (
	"go-skeleton/pkg/auth"
	"go-skeleton/pkg/common"
	"go-skeleton/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusOK, utils.JsonCodeError(common.TokenExistError))
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, utils.JsonCodeError(common.TokenTypeWrongError))
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, utils.JsonCodeError(common.TokenTypeWrongError))
			c.Abort()
			return
		}

		key, err := auth.CheckToken(checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, utils.JsonCodeError(err))
			c.Abort()
			return
		}
		c.Set("uid", key.Id)
		c.Next()
	}
}
