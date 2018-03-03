package wps

import (
	"fmt"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mailer struct {
	log    api.Logger
	config *api.Config
}

func NewMailer(log api.Logger, cfg *api.Config) Mailer {
	return Mailer{
		log:    log,
		config: cfg,
	}
}

func (m *Mailer) Send(body string) {
	to := []string{}
	for _, u := range m.config.Guests {
		to = append(to, u.Email)
	}

	from := fmt.Sprintf("%s <%s>", m.config.Admin.Name, m.config.Admin.Email)

	e := &email.Email{
		To:      to,
		From:    from,
		Subject: m.config.Subject,
		Text:    []byte(body),
	}

	err := e.Send(
		m.config.Smtp.Server+":"+m.config.Smtp.Port,
		smtp.PlainAuth(
			"",
			m.config.Admin.Email,
			m.config.Admin.Password,
			m.config.Smtp.Server,
		),
	)

	if err != nil {
		m.log.Error(fmt.Sprintf("Error sending email, [ %s ]\n", err))
		panic(err)
	}
}
