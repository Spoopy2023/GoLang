package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("The name cannot be empty!")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Goodbye(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("The name cannot be empty!")
	}
	// Return a goodbye message that embeds the name in a message.
	message := fmt.Sprintf("Goodbye, %v. See you soon!", name)
	return message, nil
}
