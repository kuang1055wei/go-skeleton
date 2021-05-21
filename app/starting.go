package app

import "gin-test/pkg/config"

func StartOn() {
	//本地无需定时任务，如果需要打开即可
	if config.Conf.ServerConfig.AppMode == "debug" {
		return
	}

	// 开启定时任务
	startSchedule()
}
