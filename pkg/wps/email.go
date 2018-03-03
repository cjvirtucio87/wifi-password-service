package wps

import (
	"fmt"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func EmailPassword(cfg *api.Config, passwd string) {
	to := []string{}
	for _, u := range cfg.Guests {
		to = append(to, u.Email)
	}

	from := fmt.Sprintf("%s <%s>", cfg.Admin.Name, cfg.Admin.Email)

	e := &email.Email{
		To:      to,
		From:    from,
		Subject: cfg.Subject,
		Text:    []byte(fmt.Sprintf("The password for the week is [ %s ]\n\nThank you!\n", passwd)),
	}

	err := e.Send(
		cfg.Smtp.Server+":"+cfg.Smtp.Port,
		smtp.PlainAuth(
			"",
			cfg.Admin.Email,
			cfg.Admin.Password,
			cfg.Smtp.Server,
		),
	)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error sending email, [ %s ]\n", err))
	}
}
