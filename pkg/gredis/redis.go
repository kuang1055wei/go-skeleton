package gredis

//https://github.com/go-redis/redis
import (
	"context"
	"fmt"
	"gin-test/pkg/config"
	"log"

	"github.com/go-redis/redis/v8"
)

//context
var Ctx = context.Background()

//redisclient
var Client *redis.Client

func InitRedis() error {

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", config.Conf.RedisConfig.Host),
		Password: config.Conf.RedisConfig.Password, // no password set
		DB:       0,                               // use default DB

		//PoolSize:     10,
		MinIdleConns: 10,
		IdleTimeout:  config.Conf.RedisConfig.IdleTimeout, //关闭空闲连接时间
	})
	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Println(err)
	}
	return nil
}
