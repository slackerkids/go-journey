package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	// if no name given, return an error with a message
	if name == "" {
		return "", errors.New("empty message")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
