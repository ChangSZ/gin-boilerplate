package admin

import (
	"encoding/json"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/api"
	"github.com/ChangSZ/gin-boilerplate/internal/code"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/core"
	"github.com/ChangSZ/gin-boilerplate/internal/pkg/password"
	"github.com/ChangSZ/gin-boilerplate/internal/repository/redis"
	"github.com/ChangSZ/gin-boilerplate/internal/services/admin"
)

type detailResponse struct {
	Username string                 `json:"username"` // 用户名
	Nickname string                 `json:"nickname"` // 昵称
	Mobile   string                 `json:"mobile"`   // 手机号
	Menu     []admin.ListMyMenuData `json:"menu"`     // 菜单栏
}

// Detail 管理员详情
// @Summary 管理员详情
// @Description 管理员详情
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/info [get]
// @Security LoginToken
func (h *handler) Detail(ctx *gin.Context) {
	res := new(detailResponse)

	searchOneData := new(admin.SearchOneData)
	searchOneData.Id = core.SessionUserInfo(ctx).UserID
	searchOneData.IsUsed = 1

	info, err := h.service.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminDetailError, err)
		return
	}

	menuCacheData, err := redis.Cache().Get(ctx, configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(searchOneData.Id)+":menu")
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminDetailError, err)
		return
	}

	var menuData []admin.ListMyMenuData
	err = json.Unmarshal([]byte(menuCacheData), &menuData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminDetailError, err)
		return
	}

	res.Username = info.Username
	res.Nickname = info.Nickname
	res.Mobile = info.Mobile
	res.Menu = menuData
	api.ResponseOK(ctx, res)
}
