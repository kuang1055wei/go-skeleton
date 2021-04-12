package router

import (
	"gin-test/controller/article"
	"gin-test/controller/site"
	"gin-test/controller/test"

	"github.com/gin-gonic/gin"
)

// func RegisterRoutes() *gin.Engine {
// 	router := gin.Default()

// 	return router
// }

func LoadDefault(e *gin.Engine) {
	e.GET("/", test.Index)
	e.GET("/helloWord", test.HelloWord)
	e.GET("/testJson", test.TestJson)

	e.GET("/site/index", site.Index)
	e.GET("/site/hello", site.Hello)
	e.POST("/site/login", site.Login)

	e.GET("/article/info", article.GetArticle)
	e.GET("/article/list", article.GetArticleList)
}
