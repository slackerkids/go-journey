package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error)  {
	// if no name given, return an error with a message
	if name == "" {
		return "", errors.New("empty message")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}