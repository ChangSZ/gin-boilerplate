package router

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ChangSZ/gin-boilerplate/assets"
	"github.com/ChangSZ/gin-boilerplate/internal/api"
	"github.com/ChangSZ/gin-boilerplate/internal/middleware"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/cron"
	"github.com/ChangSZ/gin-boilerplate/pkg/color"
	"github.com/ChangSZ/gin-boilerplate/pkg/env"
	"github.com/ChangSZ/gin-boilerplate/pkg/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const _UI = `
██████  ██ ███    ██       ██████   ██████  ██ ██      ███████ ██████  ██████  ██       █████  ████████ ███████ 
██       ██ ████   ██       ██   ██ ██    ██ ██ ██      ██      ██   ██ ██   ██ ██      ██   ██    ██    ██      
██   ███ ██ ██ ██  ██ █████ ██████  ██    ██ ██ ██      █████   ██████  ██████  ██      ███████    ██    █████   
██    ██ ██ ██  ██ ██       ██   ██ ██    ██ ██ ██      ██      ██   ██ ██      ██      ██   ██    ██    ██      
 ██████  ██ ██   ████       ██████   ██████  ██ ███████ ███████ ██   ██ ██      ███████ ██   ██    ██    ███████ 
`

func RoutersInit(cronServer cron.Server) *gin.Engine {
	if env.Active().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	eng := gin.New()
	eng.Use(
		middleware.Rate(),
		middleware.Metrics(),
		// middleware.AlertNotify(),  // 告警这个还有BUG没修复，感兴趣可以自己试着修下
		kgin.Middlewares(tracing.Server(), middleware.Logging(log.GetLoggerWithTrace()), middleware.AddTraceCtx),
	)

	fmt.Println(color.Blue(_UI))

	eng.StaticFS("assets", http.FS(assets.Bootstrap))
	eng.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))

	// 设置 Render 路由
	setRenderRouter(eng)

	// 设置 API 路由
	setApiRouter(eng, cronServer)

	// 设置 Socket 路由
	setSocketRouter(eng)

	system := eng.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx *gin.Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Request.Host,
				Status:      "ok",
			}
			api.ResponseOK(ctx, resp)
		})
	}

	var enablePProf = true
	if enablePProf {
		if !env.Active().IsPro() {
			pprof.Register(eng) // register pprof to gin
		}
	}

	if !env.Active().IsPro() {
		eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // register swagger
	}

	var enablePrometheus = true
	if enablePrometheus {
		eng.GET("/metrics", gin.WrapH(promhttp.Handler())) // register prometheus
	}
	return eng
}
