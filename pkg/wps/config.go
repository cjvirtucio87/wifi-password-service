package wps

import (
	"encoding/json"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"os/user"
	"path"
)

func ReadConfig(r api.Reader, u *user.User, filename string) api.Config {
	var cfg api.Config

	raw := r.Read(path.Join(u.HomeDir, ".wps", filename))

	json.Unmarshal(raw, &cfg)

	return cfg
}
