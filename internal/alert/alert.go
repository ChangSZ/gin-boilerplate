package alert

import (
	"github.com/ChangSZ/gin-boilerplate/configs"
	"github.com/ChangSZ/gin-boilerplate/internal/proposal"
	"github.com/ChangSZ/gin-boilerplate/pkg/errors"
	"github.com/ChangSZ/gin-boilerplate/pkg/log"
	"github.com/ChangSZ/gin-boilerplate/pkg/mail"
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

		options := &mail.Options{
			MailHost: cfg.Host,
			MailPort: cfg.Port,
			MailUser: cfg.User,
			MailPass: cfg.Pass,
			MailTo:   cfg.To,
			Subject:  subject,
			Body:     body,
		}
		if err := mail.Send(options); err != nil {
			log.Error("发送告警通知邮件失败: ", errors.WithStack(err))
		}
		return
	}
}
