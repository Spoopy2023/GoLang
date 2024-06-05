// Define the main package
package main

// add inports

import (
	"fmt"

	"rsc.io/quote"
)

// define the main function
func main() {
	fmt.Println(quote.Go()) // print the go quote imported from the quote package (rsc.io/quote)
}
