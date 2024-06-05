package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("The name cannot be empty!")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Goodbye(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("The name cannot be empty!")
	}
	//Return a goodbye message that embeds the name in a message.
	message := fmt.Sprintf("Goodbye, %v. See you soon!", name)
	return message, nil
}
