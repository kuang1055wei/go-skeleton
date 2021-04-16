package main

import (
	"fmt"
	"gin-test/logger"
	"gin-test/model"
	"gin-test/pkg/gredis"
	"gin-test/router"
	"gin-test/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//t := time.Now()
		//fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		//c.Set("request", "中间件")
		//status := c.Writer.Status()
		//fmt.Println("中间件执行完毕", status)
		//t2 := time.Since(t)
		//fmt.Println("time:", t2)
	}

}

func main() {
	//初始化配置文件
	utils.InitConfig()
	//redis
	_ = gredis.Setup()

	r := gin.New()
	gin.SetMode("debug")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//r.Use(MiddleWare())
	//r.LoadHTMLGlob("view/**/*")

	//自己用zaplog
	////logger
	//	"github.com/gin-contrib/zap"
	//logger, _ := zap.NewProduction()
	////替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	//zap.ReplaceGlobals(logger)
	//// Add a ginzap middleware, which:
	////   - Logs all requests, like a combined access and error log.
	////   - Logs to stdout.
	////   - RFC3339 with UTC time format.
	//r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	//
	//// Logs all panic to error log
	////   - stack means whether output the stack info.
	//r.Use(ginzap.RecoveryWithZap(logger, true))
	////zaplog结束

	//使用自定义的zaplog,且日志归档到文件
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//log结束

	//logger end
	//加载路由
	router.LoadDefault(r)
	//数据库初始化
	model.InitDb()
	_ = r.Run(":8000")
}
