package alert

import (
	"strings"

	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/mail"

	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/proposal"
	"github.com/ChangSZ/gin-boilerplate/pkg/errors"
)

// NotifyHandler 告警通知
func NotifyHandler() func(msg *proposal.AlertMessage) {
	return func(msg *proposal.AlertMessage) {
		cfg := configs.Get().Mail
		if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
			log.Error("Mail config error")
			return
		}

		subject, body, err := newHTMLEmail(
			msg.Method,
			msg.HOST,
			msg.URI,
			msg.TraceID,
			msg.ErrorMessage,
			msg.ErrorStack,
		)
		if err != nil {
			log.Error("email template error: ", err)
			return
		}

		client, err := mail.Init(
			mail.WithUser(cfg.User),
			mail.WithPwd(cfg.Pass),
			mail.WithHost(cfg.Host),
			mail.WithPort(cfg.Port))
		if err != nil {
			log.Error("邮件client初始化失败: ", err)
		}
		if err := client.SetTo(strings.Split(cfg.To, ",")).
			SetSubject(subject).
			SetBody(body).
			Send(); err != nil {
			log.Error("发送告警通知邮件失败: ", errors.WithStack(err))
		}
	}
}
