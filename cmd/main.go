package main

import (
	"github.com/cjvirtucio87/wifi-password-service/internal/wps"
	"log"
)

func main() {
	cfg := wps.ConfigJson()
	passwd := wps.Password(cfg)

	log.Println("Emailing users.")
	wps.EmailPassword(cfg, passwd)

	log.Println("Done!")
}
