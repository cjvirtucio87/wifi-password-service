package wps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Sender struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Smtp struct {
	Server string `json:"server,omitempty"`
	Port   string `json:"port,omitempty"`
}

type Config struct {
	Users   []User   `json:"users,omitempty"`
	Words   []string `json:"words,omitempty"`
	Sender  Sender   `json:"sender,omitempty"`
	Smtp    Smtp     `json:"smtp,omitempty"`
	Body    string   `json:"body,omitempty"`
	Subject string   `json:"subject,omitempty"`
}

func ConfigJson() Config {
	var cfg Config

	usr, err := user.Current()

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to retrieve user due to error, [%s]", err))
	}

	raw, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".wps", "config.json"))

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read file, [%s]", err))
	}

	json.Unmarshal(raw, &cfg)

	return cfg
}
