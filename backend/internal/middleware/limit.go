package middleware

import (
	"backend/pkg/ginx"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

const (
	_rate    = 3
	_cap     = 10
	_timeout = 500
)

func Limiter() gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		l, _ := limiters.LoadOrStore(ip, rate.NewLimiter(_rate, _cap))
		c, cancel := context.WithTimeout(ctx, time.Millisecond*_timeout)
		defer cancel()
		if err := l.(*rate.Limiter).Wait(c); err != nil {
			ginx.BeyondRateLimit(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
