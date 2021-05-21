package app

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

func startSchedule() {
	c := cron.New()

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

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		zap.L().Error("添加计划任务失败", zap.NamedError("error:", err))
	}
}
