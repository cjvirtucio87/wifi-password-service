package wps

import (
	"math/rand"
	"strings"
	"time"
)

func Password(cfg Config) string {
	rand.Seed(time.Now().UnixNano())
	wordCount := len(cfg.Words)
	selected := []string{}

	for i := 0; i < 4; i++ {
		randWord := cfg.Words[rand.Intn(wordCount)]
		selected = append(selected, randWord)
	}

	return strings.Join(selected, " ")
}
