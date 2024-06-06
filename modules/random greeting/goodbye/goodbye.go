package main

import (
	"fmt"

	"github.com/spoopy2023/greetings"
)

func main() {
	message := greetings.Goodbye("Spoopy2023")
	fmt.Println(message)
}
