package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/ChangSZ/gin-boilerplate/configs"
)

// Rate 限流
func Rate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := rate.NewLimiter(rate.Every(time.Second*1), configs.MaxRequestsPerSecond)
		if !limiter.Allow() {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		ctx.Next()
	}
}
