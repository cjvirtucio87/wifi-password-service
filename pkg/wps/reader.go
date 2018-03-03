package wps

import (
	"fmt"
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"io/ioutil"
)

type ConfigReader struct {
	log api.Logger
}

func NewConfigReader(log api.Logger) api.Reader {
	return ConfigReader{
		log: log,
	}
}

func (r ConfigReader) Read(filename string) []byte {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		r.log.Error((fmt.Sprintf("Failed to read file, [%s]", err)))
		panic(err)
	}

	return raw
}
