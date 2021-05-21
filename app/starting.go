package app

import "gin-test/pkg/config"

func StartOn() {
	if config.Conf.ServerConfig.AppMode == "debug" {
		return
	}

	// 开启定时任务
	startSchedule()
}
