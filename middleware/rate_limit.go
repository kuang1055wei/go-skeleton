package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go.uber.org/zap"
	"net/http"
	"time"
)

//限流中间件 RateLimitMiddleware(time.Second,1),表示1秒钟一个请求
func RateLimitMiddleware(fillInterval time.Duration , cap int64)func(c *gin.Context)  {
	//创建指定填充速率和容量大小的令牌桶
	bucket := ratelimit.NewBucket(fillInterval , cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		zap.L().Info(fmt.Sprintf("bucket:%d\n" , bucket.Rate()))
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusOK , gin.H{
				"code":400,
				"message":"访问太快了，请慢一点",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
