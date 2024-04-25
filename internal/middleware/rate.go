package middleware

import (
	"net/http"
	"time"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/code"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Rate 限流
func Rate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := rate.NewLimiter(rate.Every(time.Second*1), configs.MaxRequestsPerSecond)
		if !limiter.Allow() {
			core.AbortWithError(ctx, core.Error(
				http.StatusTooManyRequests,
				code.TooManyRequests,
				code.Text(code.TooManyRequests)),
			)
			return
		}

		ctx.Next()
	}
}
