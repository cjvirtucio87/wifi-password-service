package main

import (
	"fmt"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"github.com/cjvirtucio87/wifi-password-service/pkg/wps"
	"github.com/urfave/cli"
	"os"
	"time"
)

func configFile(net string) string {
	var filename string

	switch net {
	case "5g":
		filename = "config-5g.json"
		break
	default:
		filename = "config.json"
		break
	}

	return filename
}

func elapsed(log api.Logger, t time.Time) {
	log.Info("Done!")
	log.Info(fmt.Sprintf("Time elapsed: %s", time.Since(t)))
}

func do(c *cli.Context) {
	log := wps.NewWpsLogger(c.String("log"))
	start := time.Now()
	defer elapsed(log, start)

	r := wps.NewConfigReader(&log)
	usr := wps.User(&log)
	cfg := wps.ReadConfig(r, usr, configFile(c.String("net")))

	m := wps.NewMailer(&log, &cfg)

	passwd := wps.Password(&cfg)

	log.Info("Emailing users.")
	m.Send(fmt.Sprintf("The wifi password for the week is [%s]. Please update your wifi connection settings, now.\nThank you!\n", passwd))
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "net",
			Value: "default",
			Usage: "router network to use, i.e. whether it's the regular one, or the 5g one (defaults to regular)",
		},
		cli.StringFlag{
			Name:  "log",
			Value: "info",
			Usage: "log level for the app",
		},
	}

	app.Action = do

	err := app.Run(os.Args)

	if err != nil {
		fmt.Printf("Failed to run app due to error, [%s]", err)
		panic(err)
	}
}
