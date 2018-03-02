package main

import (
	"github.com/cjvirtucio87/wifi-password-service/internal/wps"
	"log"
)

func main() {
	cfg := wps.ConfigJson()
	users := cfg.Users
	passwd := wps.Password(cfg)

	log.Println("Emailing users.")
	for _, u := range users {
		wps.EmailPassword(u, passwd)
	}

	log.Println("Done!")
}
