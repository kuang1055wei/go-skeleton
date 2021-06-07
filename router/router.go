package router

import (
	"go-skeleton/controller/site"

	"github.com/gin-gonic/gin"
)

func LoadDefault(e *gin.Engine) {

	e.GET("/site/index", site.Index)
	e.GET("/site/hello", site.Hello)
	e.POST("/site/login", site.Login)
	e.POST("/site/refreshToken", site.RefreshAccessToken)

}
