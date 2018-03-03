package wps

import (
	"fmt"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"os/user"
)

func User(log api.Logger) *user.User {
	usr, err := user.Current()

	if err != nil {
		log.Error(fmt.Sprintf("Failed to retrieve user due to error, [%s]", err))
		panic(err)
	}

	return usr
}
