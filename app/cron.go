package app

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Logger struct {
}

func startSchedule() {
	//SkipIfStillRunning为前面任务没执行完，则跳过当前任务
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))

	//c.AddFunc("1/2 * * * * ?", func() {
	//	fmt.Println("2s一次:", carbon.Now().ToDateTimeString())
	//	time.Sleep(time.Second * 10)
	//})

	//加锁保证多机器只执行一次
	//addCronFunc(c, "1/5 * * * * ?", func() {
	//	nx, _ := gredis.GetRedis().SetNX(context.TODO(), "test_cron", 1, time.Minute).Result()
	//	fmt.Println("锁结果", nx)
	//	if nx {
	//		defer func() {
	//			fmt.Println("15s后，执行成功释放锁", nx)
	//			gredis.GetRedis().Del(context.TODO(), "test_cron")
	//		}()
	//		fmt.Println("计划任务执行")
	//	} else {
	//		fmt.Println("----计划任务不执行----")
	//	}
	//	return
	//})

	//addCronFunc(c , "0/5 * * * * ?" , func() {
	//	course := services.CourseExchangeCodeService
	//	_ = course.GenerateCodeFromRedis()
	//})

	//addCronFunc(c, "@every 30m", func() {
	//})

	// Generate sitemap
	//addCronFunc(c, "0 0 4 ? * *", func() {
	//})

	c.Start()
}

//如果是负载均衡部署，cmd这里添加的计划任务都应该用redis或者etcd等加锁，防止任务重复执行
func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	_, err := c.AddFunc(sepc, cmd)
	if err != nil {
		zap.L().Error("添加计划任务失败", zap.NamedError("error:", err))
	}
}
