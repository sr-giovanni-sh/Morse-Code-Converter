package service

import (
	"Morse-Code-Converter/pkg/morse"
	"errors"
)

// isMorse checks if the data is in Morse code format.
func isMorse(data string) bool {
	for _, v := range data {
		if v != '.' && v != ' ' && v != '-' && v != '/' && v != '\n' && v != '\r' {
			return false
		}
	}
	return true
}

// AutoConvert automatically converts data between text and Morse code.
func AutoConvert(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty input")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}
