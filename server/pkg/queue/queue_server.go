package queue

import (
	"fmt"
	"go-skeleton/pkg/config"
	"go-skeleton/task"

	"github.com/RichardKnop/machinery/v1"
	queueConfig "github.com/RichardKnop/machinery/v1/config"
)

var server *machinery.Server

func GetServer() *machinery.Server {
	return server
}

func startServer() (*machinery.Server, error) {
	//redis://password@localhost:6379
	//redis://localhost:6379
	var addr string
	if config.Conf.RedisConfig.Password == "" {
		addr = fmt.Sprintf("redis://%s", config.Conf.RedisConfig.Host)
	} else {
		addr = fmt.Sprintf("redis://%s@%s", config.Conf.RedisConfig.Password, config.Conf.RedisConfig.Host)
	}

	conf := &queueConfig.Config{
		Broker:          addr,
		DefaultQueue:    "default_queue",
		ResultBackend:   addr, //结果保存的地方
		ResultsExpireIn: 60,
		Redis: &queueConfig.RedisConfig{
			MaxIdle:     config.Conf.RedisConfig.MaxIdle,
			MaxActive:   config.Conf.RedisConfig.MaxActive,
			IdleTimeout: int(config.Conf.RedisConfig.IdleTimeout),
		},
	}

	ser, err := machinery.NewServer(conf)
	if err != nil {
		return nil, err
	}

	return ser, nil
}

func InitQueue() error {

	var err error
	server, err = startServer()
	if err != nil {
		return err
	}
	//注册task
	//server.RegisterTasks(map[string]interface{}{
	//	"add":      Add,
	//	"multiply": Multiply,
	//})
	_ = server.RegisterTask("add", task.Add)
	_ = server.RegisterTask("panic", task.PanicTask)
	_ = server.RegisterTask("error", task.ErrorTask)

	//起一个协程执行注册worker执行任务
	go func(s *machinery.Server) {
		worker := s.NewWorker("default_queue_work", 10)
		//处理错误handle
		//worker.SetErrorHandler(func(err error) {
		//
		//})
		_ = worker.Launch()
	}(server)

	return nil
}
