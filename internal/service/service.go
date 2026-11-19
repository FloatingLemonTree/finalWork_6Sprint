package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	value := strings.TrimSpace(input)
	if value == "" {
		return "", errors.New("пустой ввод")
	}

	isMorse := strings.Trim(value, ".- ") == ""
	if !isMorse {
		return morse.ToMorse(value), nil
	}

	return morse.ToText(value), nil
}
