package router

import (
	"fmt"
	"gin-test/controller/article"
	"gin-test/controller/site"
	"gin-test/controller/test"
	"gin-test/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// func RegisterRoutes() *gin.Engine {
// 	router := gin.Default()

// 	return router
// }
// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		//设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}

}
func LoadDefault(e *gin.Engine) {
	e.GET("/", test.Index)
	e.GET("/helloWord", test.HelloWord)
	e.GET("/testJson", test.TestJson)

	e.GET("/site/index", site.Index)
	e.GET("/site/hello", site.Hello)
	e.POST("/site/login", site.Login)

	e.GET("/article/info", article.GetArticle)
	e.GET("/article/list", article.GetArticleList)
	e.GET("/article/myHttp", article.MyHttp)
	e.GET("/article/myChan", article.MyChan)
	e.POST("/article/edit", article.EditArticle)
	e.GET("/article/search", article.SearchArticle)
	e.POST("/article/uploadImg" , article.UploadImg)
	//路由组
	v1 := e.Group("v1").Use(middleware.JwtToken())
	{
		v1.GET("/article/info", article.GetArticle)
		v1.GET("/article/list", article.GetArticleList)
	}
}
