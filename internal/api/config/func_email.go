package config

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/mail"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/api"
	"github.com/ChangSZ/gin-boilerplate/internal/code"
	"github.com/ChangSZ/gin-boilerplate/pkg/env"
	"github.com/ChangSZ/gin-boilerplate/pkg/validator"
)

type emailRequest struct {
	Host string `form:"host"` // 邮箱服务器
	Port string `form:"port"` // 端口
	User string `form:"user"` // 发件人邮箱
	Pass string `form:"pass"` // 发件人密码
	To   string `form:"to"`   // 收件人邮箱地址，多个用,分割
}

type emailResponse struct {
	Email string `json:"email"` // 邮箱地址
}

// Email 修改邮件配置
// @Summary 修改邮件配置
// @Description 修改邮件配置
// @Tags API.config
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param host formData string true "邮箱服务器"
// @Param port formData string true "端口"
// @Param user formData string true "发件人邮箱"
// @Param pass formData string true "发件人密码"
// @Param to formData string true "收件人邮箱地址，多个用,分割"
// @Success 200 {object} emailResponse
// @Failure 400 {object} code.Failure
// @Router /api/config/email [patch]
// @Security LoginToken
func (h *handler) Email(ctx *gin.Context) {
	req := new(emailRequest)
	res := new(emailResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	client, err := mail.Init(
		mail.WithUser(req.User),
		mail.WithPwd(req.Pass),
		mail.WithHost(req.Host),
		mail.WithPort(cast.ToInt(req.Port)))
	if err != nil {
		err := fmt.Errorf("邮件client初始化失败: %w", err)
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.SendEmailError, err)
	}

	if err := client.SetTo(strings.Split(req.To, ",")).
		SetSubject(fmt.Sprintf("%s[%s] 邮箱告警人调整通知。", configs.ProjectName, env.Active().Value())).
		SetBody(fmt.Sprintf("%s[%s] 已添加您为系统告警通知人。", configs.ProjectName, env.Active().Value())).
		Send(); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.SendEmailError, err)
	}

	viper.Set("mail.host", req.Host)
	viper.Set("mail.port", cast.ToInt(req.Port))
	viper.Set("mail.user", req.User)
	viper.Set("mail.pass", req.Pass)
	viper.Set("mail.to", req.To)

	err = viper.WriteConfig()
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.WriteConfigError, err)
		return
	}

	res.Email = req.To
	api.ResponseOK(ctx, res)
}
