package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 繼承Limiter結構體
type MethodLimiter struct {
	*Limiter
}

// 創建一個新的MothodLimiter結構體
func NewMMethodLimiter() LimiterIface {
	//初始化Limiter中的limiterBuckets欄位，創建一個map
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	//example: /api/resource?id=123
	uri := c.Request.RequestURI
	//查詢?號的index位置
	index := strings.Index(uri, "?")
	//如果沒有查到是-1，則返回整個uri
	if index == -1 {
		return uri
	}
	//如果查到了，則返回查詢到的index之前的所有字串 :/api/resource
	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}
	return l
}
