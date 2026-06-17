package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) (error, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return errors.New("Empty string"), false
	}

	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return nil, false
		}
	}

	return nil, true
}

func ConvertText(text string) string {
	err, check := isMorse(text)
	if err != nil {
		log.Fatal(err)
	}
	if check {
		return morse.ToText(text)
	}

	return morse.ToMorse(text)
}
