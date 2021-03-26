package main

import (
	"fmt"

	"github.com/dubiZA/golang_learn/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
