package gredis

//https://github.com/gomodule/redigo
//import (
//"encoding/json"
//"gin-test/utils"
//"time"
//
//"github.com/gomodule/redigo/redis"
//)
//
//var RedisConn *redis.Pool
//
//// Setup Initialize the Redis instance
//func Setup() error {
//	RedisConn = &redis.Pool{
//		MaxIdle:     utils.Conf.RedisConfig.MaxIdle,
//		MaxActive:   utils.Conf.RedisConfig.MaxActive,
//		IdleTimeout: utils.Conf.RedisConfig.IdleTimeout,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", utils.Conf.RedisConfig.Host)
//			if err != nil {
//				return nil, err
//			}
//			if utils.Conf.RedisConfig.Password != "" {
//				if _, err := c.Do("AUTH", utils.Conf.RedisConfig.Password); err != nil {
//					_ = c.Close()
//					return nil, err
//				}
//			}
//			return c, err
//		},
//		TestOnBorrow: func(c redis.Conn, t time.Time) error {
//			_, err := c.Do("PING")
//			return err
//		},
//	}
//
//	return nil
//}
//
//// Set a key/value
//func Set(key string, data interface{}, time int) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	value, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	_, err = conn.Do("SET", key, value, "ex", time)
//	if err != nil {
//		return err
//	}
//
//	//_, err = conn.Do("EXPIRE", key, time)
//	//if err != nil {
//	//	return err
//	//}
//
//	return nil
//}
//
////setnx ok等于拿到锁，
//func SetNx(key string, data interface{}, time int) (bool, error) {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	value, err := json.Marshal(data)
//	if err != nil {
//		return false, err
//	}
//
//	reply, err := redis.String(conn.Do("SET", key, value, "nx", "ex", time))
//	if err != nil {
//		return false, err
//	}
//
//	return reply == "OK", err
//}
//
//// Exists check a key
//func Exists(key string) bool {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	exists, err := redis.Bool(conn.Do("EXISTS", key))
//	if err != nil {
//		return false
//	}
//
//	return exists
//}
//
//// Get get a key
//func Get(key string) ([]byte, error) {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	reply, err := redis.Bytes(conn.Do("GET", key))
//	if err != nil {
//		return nil, err
//	}
//
//	return reply, nil
//}
//
////删除key1,key2,keyN...
//func Del(key ...interface{}) (int64, error) {
//	conn := RedisConn.Get()
//	defer conn.Close()
//	return redis.Int64(conn.Do("DEL", key...))
//}
//
//// LikeDeletes batch delete
//func LikeDeletes(key string) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
//	if err != nil {
//		return err
//	}
//
//	for _, key := range keys {
//		_, err = Del(key)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
