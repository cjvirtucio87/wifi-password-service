package main

import (
	"github.com/cjvirtucio87/wifi-password-service/pkg/wps"
)

func main() {
	log := wps.NewWpsLogger("debug")

	r := wps.NewConfigReader(log)
	usr := wps.User(log)
	cfg := wps.ReadConfig(r, usr, "config.json")

	passwd := wps.Password(&cfg)

	log.Info("Emailing users.")
	wps.EmailPassword(&cfg, passwd)

	log.Info("Done!")
}
