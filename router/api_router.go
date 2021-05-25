package router

import (
	"gin-test/controller/api"
	"gin-test/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(e *gin.Engine) {
	art := api.ArticleController{}
	captch := api.CaptchaController{}
	//路由组
	//apiRouter := e.Group("api").Use(middleware.RateLimitMiddleware(time.Second, 100)).Use(middleware.JwtToken())
	apiRouter := e.Group("api").Use(middleware.RateLimitMiddleware(time.Second, 100))
	{
		apiRouter.GET("/article/getArticleById", art.GetArticleById)
		apiRouter.GET("/article/getArticleByCache", art.GetArticleByCache)
		apiRouter.GET("/article/create", art.Create)
		apiRouter.GET("/article/info", art.GetArticle)
		apiRouter.GET("/article/list", art.GetArticleList)
		apiRouter.GET("/article/myHttp", art.MyHttp)
		apiRouter.GET("/article/myChan", art.MyChan)
		apiRouter.POST("/article/edit", art.EditArticle)
		apiRouter.GET("/article/search", art.SearchArticle)
		apiRouter.POST("/article/uploadImg", art.UploadImg)
		apiRouter.GET("/article/viperTest", art.ViperTest)
		apiRouter.GET("/article/myChan2", art.MyChan2)
		apiRouter.GET("/article/queue", art.TestQueue)
		apiRouter.GET("/article/testPanic", art.TestPanic)
		apiRouter.GET("/captcha", captch.GenerateCaptcha)
		apiRouter.GET("/verifyCaptcha", captch.VerifyCaptcha)
		apiRouter.GET("/redisLock", art.TryRedisLock)
	}
}
