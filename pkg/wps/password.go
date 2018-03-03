package wps

import (
	"github.com/cjvirtucio87/wifi-password-service/pkg/api"
	"math/rand"
	"strings"
	"time"
)

func Password(cfg *api.Config) string {
	rand.Seed(time.Now().UnixNano())
	wordCount := len(cfg.Words)
	selected := []string{}

	for i := 0; i < 4; i++ {
		randWord := cfg.Words[rand.Intn(wordCount)]
		selected = append(selected, randWord)
	}

	return strings.Join(selected, " ")
}
