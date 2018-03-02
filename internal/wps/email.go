package wps

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func EmailPassword(cfg Config, passwd string) {
	to := []string{}
	for _, u := range cfg.Users {
		to = append(to, u.Email)
	}

	from := fmt.Sprintf("%s <%s>", cfg.Sender.Name, cfg.Sender.Email)

	e := &email.Email{
		To:      to,
		From:    from,
		Subject: cfg.Subject,
		Text:    []byte(fmt.Sprintf("The password for the week is [ %s ]\n\nThank you!\n", passwd)),
	}

	log.Println(cfg.Sender.Email)
	log.Println(cfg.Sender.Password)

	err := e.Send(
		cfg.Smtp.Server+":"+cfg.Smtp.Port,
		smtp.PlainAuth(
			"",
			cfg.Sender.Email,
			cfg.Sender.Password,
			cfg.Smtp.Server,
		),
	)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error sending email, [ %s ]\n", err))
	}
}
