package gredis

//https://github.com/go-redis/redis
import (
	"context"
	"fmt"
	"go-skeleton/pkg/config"

	"github.com/go-redis/redis/v8"
)

//redisclient
var client *redis.Client

func GetRedis() *redis.Client {
	return client
}

func InitRedis() error {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", config.Conf.RedisConfig.Host),
		Password: config.Conf.RedisConfig.Password, // no password set
		DB:       0,                                // use default DB

		//PoolSize:     10,
		MinIdleConns: 10,
		IdleTimeout:  config.Conf.RedisConfig.IdleTimeout, //关闭空闲连接时间
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return err
	}
	return nil
}
