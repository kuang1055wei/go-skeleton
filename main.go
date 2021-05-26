package main

import (
	"context"
	"fmt"
	"gin-test/app"
	"gin-test/logger"
	"gin-test/middleware"
	"gin-test/pkg/config"
	"gin-test/pkg/gredis"
	"gin-test/pkg/simpleDb"
	"gin-test/pkg/upload"
	"gin-test/router"
	"gin-test/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	if err := gredis.InitRedis(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		os.Exit(0)
		return
	}

	//数据库初始化
	if err := simpleDb.InitDb(); err != nil {
		fmt.Printf("连接数据库失败，请检查参数:%v\n", err)
		os.Exit(0)
		return
	}

	//翻译
	_ = utils.InitTrans("zh")

	//队列任务初始化，如果要使用队列打开注释即可，task目录为处理方法，在InitQueue中注册即可
	//if err := queue.InitQueue(); err != nil {
	//	fmt.Printf("队列初始化失败:%v\n", err)
	//	os.Exit(0)
	//	return
	//}
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
	r.Use(middleware.Cors())

	//加载路由
	router.LoadDefault(r)
	router.LoadAdminRouter(r)
	router.LoadApiRouter(r)
	return r
}

func main() {
	app.StartOn()

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

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")

	// 默认endless服务器会监听下列信号：
	//	"github.com/fvbock/endless"
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//if err := endless.ListenAndServe(address , handler); err!=nil{
	//	log.Fatalf("listen: %s\n", err)
	//}
}
