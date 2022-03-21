package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		&Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

//MethodLimiter实现LimiterIface接口
func (m MethodLimiter)Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri,"?")
	if index == -1{
		return uri
	}
	return uri[:index]
}

func (m MethodLimiter)GetBucket(key string) (*ratelimit.Bucket,bool) {
	bucket,ok := m.limiterBuckets[key]
	return bucket,ok
}

func (m MethodLimiter)AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _,rule := range rules{
		if _,ok := m.limiterBuckets[rule.Key]; !ok{
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval,rule.Capacity,rule.Quantum)
		}
	}
	return m
}


