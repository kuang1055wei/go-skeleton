package gredis

//https://github.com/go-redis/redis
import (
	"context"
	"fmt"
	"gin-test/utils"

	"github.com/go-redis/redis/v8"
)

//context
var Ctx = context.Background()

//redisclient
var RedisClient *redis.Client

func Setup() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s", utils.Conf.RedisConfig.Host),
		Password:    utils.Conf.RedisConfig.Password,    // no password set
		DB:          0,                                  // use default DB
		IdleTimeout: utils.Conf.RedisConfig.IdleTimeout, //关闭空闲连接时间
	})
	return nil
}
