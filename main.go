package main

import (
	"fmt"
	"gin-test/logger"
	"gin-test/model"
	"gin-test/pkg/configloader"
	"gin-test/pkg/gredis"
	"gin-test/router"
	"gin-test/utils"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func init() {
	//初始化配置文件
	utils.InitConfig()
	configloader.InitConfig()
	//redis
	_ = gredis.Setup()
	//翻译
	_ = utils.InitTrans("zh")

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
	//log结束

	//数据库初始化
	model.InitDb()
}

func initGin() *gin.Engine {
	gin.SetMode(utils.Conf.ServerConfig.AppMode)
	r := gin.New()
	//session
	//store, _ := redis.NewStore(
	//	10,
	//	"tcp",
	//	fmt.Sprintf("%s", utils.Conf.RedisConfig.Host),
	//	utils.Conf.RedisConfig.Password,
	//	[]byte("secret"),
	//)
	//r.Use(sessions.Sessions("mysession", store))

	//r.Use(MiddleWare())
	//r.LoadHTMLGlob("view/**/*")

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//加载路由
	router.LoadDefault(r)
	return r
}

//func main2() {
//	r := initGin()
//	_ = r.Run(fmt.Sprintf(":%s", utils.Conf.ServerConfig.HttpPort)) //:8000
//}

func main() {

	handler := initGin()
	address := fmt.Sprintf(":%s", utils.Conf.ServerConfig.HttpPort)
	readTimeout := utils.Conf.ServerConfig.ReadTimeout * time.Second
	writeTimeout := utils.Conf.ServerConfig.WriteTimeout * time.Second
	maxHeaderBytes := 1 << 20 //1048576 = 1mb

	server := &http.Server{
		Addr:           address,
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	zap.L().Info(fmt.Sprintf("Listening and serving HTTP on %s\n", address))

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
