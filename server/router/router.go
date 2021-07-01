package router

import (
	"go-skeleton/controller/site"

	_ "go-skeleton/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func LoadDefault(e *gin.Engine) {

	e.GET("/site/index", site.Index)
	e.GET("/site/hello", site.Hello)
	e.POST("/site/login", site.Login)
	e.POST("/site/refreshToken", site.RefreshAccessToken)
	e.POST("/site/register", site.Register)

	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
