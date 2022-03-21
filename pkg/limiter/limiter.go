package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//LimiterIface接口定义限流器所必须的方法
type LimiterIface interface {
	//获取限流器的键值对名称
	Key(c *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket,bool)
	//新增令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key string
	FillInterval time.Duration
	Capacity int64
	Quantum int64
}
