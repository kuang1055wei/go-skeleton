package gcaptcha

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	ctx   context.Context
	redis *redis.Client
}

func NewRedisStore(ctx context.Context, redis *redis.Client) *redisStore {
	return &redisStore{
		ctx:   ctx,
		redis: redis,
	}
}

//获取验证码的key
func (r *redisStore) getKey(id string) string {
	return "captcha:" + id
}

// 设置验证码ID的数字
func (r *redisStore) Set(id string, value string) {
	_ = r.redis.Set(r.ctx, r.getKey(id), value, time.Minute*5).Err()
}

//获取验证码ID的返回存储数字。清除表示是否必须从商店中删除验证码。
func (r *redisStore) Get(id string, clear bool) string {
	val, _ := r.redis.Get(r.ctx, r.getKey(id)).Result()
	if clear {
		_ = r.redis.Del(r.ctx, r.getKey(id)).Err()
	}
	return val
}

//直接验证验证码的答案,忽略大小写
func (r *redisStore) Verify(id, answer string, clear bool) bool {
	val := r.Get(id, clear)
	return strings.ToLower(val) == strings.ToLower(answer)
}
