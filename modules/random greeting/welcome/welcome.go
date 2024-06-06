package main

import (
	"fmt"
	"log"

	"github.com/spoopy2023/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:Welcome - Logger ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Tyler") // this will return an error due to the empty string, in this case the error message is "The name cannot be empty!"

	// If an error was returned, print it to the console and kill the program.
	if err != nil {
		log.Fatal(err)
	}
	// if no error was returned, print the message to the console.
	fmt.Println(message)
}
