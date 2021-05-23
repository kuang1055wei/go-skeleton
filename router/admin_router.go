package router

import (
	"gin-test/controller/admin"
	"gin-test/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func LoadAdminRouter(e *gin.Engine) {
	art := admin.ArticleController{}
	//路由组
	adminRouter := e.Group("admin").Use(middleware.RateLimitMiddleware(time.Second, 100)).Use(middleware.JwtToken())
	//adminRouter := e.Group("admin").Use(middleware.RateLimitMiddleware(time.Second, 100))
	{
		adminRouter.GET("/article/info", art.GetArticle)
		adminRouter.GET("/article/list", art.GetArticleList)
		adminRouter.GET("/article/myHttp", art.MyHttp)
		adminRouter.GET("/article/myChan", art.MyChan)
		adminRouter.POST("/article/edit", art.EditArticle)
		adminRouter.GET("/article/search", art.SearchArticle)
		adminRouter.POST("/article/uploadImg", art.UploadImg)
		adminRouter.GET("/article/viperTest", art.ViperTest)
		adminRouter.GET("/article/myChan2", art.MyChan2)
		adminRouter.GET("/article/queue", art.TestQueue)
	}
}
