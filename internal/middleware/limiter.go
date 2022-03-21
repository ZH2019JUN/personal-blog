package middleware

import (
	"github.com/gin-gonic/gin"
	"myproject/pkg/app"
	"myproject/pkg/errorcode"
	"myproject/pkg/limiter"
)

//限流器,可以对指定接口进行限流控制
func RateLimiter(m limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := m.Key(c)
		bucket,ok := m.GetBucket(key)
		if ok{
			//TakeAvailable方法占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数量，若没有可用令牌，则返回0
			count := bucket.TakeAvailable(1)
			if count==0{
				app.NewResponse(c).ToErrorResponse(errorcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
