package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/api"
	"github.com/ChangSZ/gin-boilerplate/internal/code"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
	"github.com/ChangSZ/gin-boilerplate/internal/services/admin"
	"github.com/ChangSZ/gin-boilerplate/pkg/urltable"
)

// CheckRBAC 验证 RBAC 权限是否合法
func CheckRBAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Token")
		if token == "" {
			err := errors.New("header 中缺少 Token 参数")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if !redis.Cache().Exists(ctx, configs.RedisKeyPrefixLoginUser+token) {
			err := errors.New("请先登录")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.CacheGetError, err)
			ctx.Abort()
			return
		}

		if !redis.Cache().Exists(ctx, configs.RedisKeyPrefixLoginUser+token+":action") {
			err := errors.New("当前账号未配置 RBAC 权限")
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.CacheGetError, err)
			ctx.Abort()
			return
		}

		actionData, err := redis.Cache().Get(ctx, configs.RedisKeyPrefixLoginUser+token+":action")
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.CacheGetError, err)
			ctx.Abort()
			return
		}

		var actions []admin.MyActionData
		err = json.Unmarshal([]byte(actionData), &actions)
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusUnauthorized, code.AuthorizationError, err)
			ctx.Abort()
			return
		}

		if len(actions) > 0 {
			urlPath := ctx.Request.URL.Path
			method := ctx.Request.Method
			table := urltable.NewTable()
			for _, v := range actions {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(method + urlPath); pattern == "" {
				err := errors.New(method + urlPath + " 未进行 RBAC 授权")
				log.WithTrace(ctx).Error(err)
				api.Response(ctx, http.StatusBadRequest, code.RBACError, err)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
