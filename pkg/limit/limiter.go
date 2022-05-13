package limit

import (
	"github/leave8080/go_package/pkg/eno"
	"net/http"
	"strings"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

var (
	defaultConfig = &Config{
		Rate:       1000,
		BucketSize: 1000,
	}
)

type Config struct {
	Rate       int // per second request，0 不限流
	BucketSize int // max size，桶内最大量
}

// 速率限制器
type RateLimiter struct {
	C            *Config
	LimiterGroup *RateGroup
}

func NewLimiter(c *Config) (rl *RateLimiter) {
	if c == nil {
		c = defaultConfig
	}
	rl = &RateLimiter{
		C: c,
		LimiterGroup: NewRateGroup(func() interface{} {
			return newLimiter(c)
		}),
	}
	return rl
}

func (r *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := strings.Split(c.Request.RequestURI, "?")[0]
		//log.Warning("key:", path[1:])
		limiter := r.LimiterGroup.Get(path[1:])
		if allow := limiter.Allow(); !allow {
			e := eno.AnalyseError(eno.ServerBusyErr)
			rsp := struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    e.Code(),
				Message: e.Message(),
			}
			c.JSON(http.StatusOK, rsp)
			c.Abort()
		}
		c.Next()
	}
}

func newLimiter(c *Config) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(c.Rate), c.BucketSize)
}
