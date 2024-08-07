package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/metrics"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/proposal"
	"github.com/ChangSZ/gin-boilerplate/pkg/env"
)

// Metrics 统计
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		ts := time.Now()
		defer func() {
			path := ctx.Request.URL.Path
			if alias := core.Alias(ctx); alias != "" {
				path = alias
			}

			metrics.RecordHandler()(&proposal.MetricsMessage{
				ProjectName: configs.ProjectName,
				Env:         env.Active().Value(),
				TraceID:     core.TraceID(ctx),
				HOST:        ctx.Request.Host,
				Path:        path,
				Method:      ctx.Request.Method,
				HTTPCode:    ctx.Writer.Status(),
				CostSeconds: time.Since(ts).Seconds(),
				IsSuccess:   !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK),
			})
		}()

		ctx.Next()
	}
}
