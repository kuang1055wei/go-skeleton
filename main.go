package main

import (
	"fmt"
	"gin-test/logger"
	"gin-test/model"
	"gin-test/pkg/config"
	"gin-test/pkg/gredis"
	"gin-test/pkg/upload"
	"gin-test/router"
	"gin-test/utils"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func init() {
	//初始化配置文件
	//utils.InitConfig()
	if err := config.InitConfig(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		os.Exit(0)
		return
	}

	//使用自定义的zaplog,且日志归档到文件
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	//log结束

	//redis
	if err := gredis.InitRedis();err!=nil{
		fmt.Printf("init redis failed, err:%v\n", err)
		os.Exit(0)
		return
	}

	//数据库初始化
	if err := model.InitDb();err!=nil{
		fmt.Printf("连接数据库失败，请检查参数:%v\n", err)
		os.Exit(0)
		return
	}

	//翻译
	_ = utils.InitTrans("zh")
}

func initGin() *gin.Engine {
	gin.SetMode(config.Conf.ServerConfig.AppMode)
	r := gin.New()
	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
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
	address := fmt.Sprintf(":%s", config.Conf.ServerConfig.HttpPort)
	readTimeout := config.Conf.ServerConfig.ReadTimeout * time.Second
	writeTimeout := config.Conf.ServerConfig.WriteTimeout * time.Second
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
