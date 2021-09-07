package main

import (
	"fmt"
	"log"

	// "rsc.io/quote"
	"example.com/greetings"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	// Get a greeting message and print it
	// message := greetings.Hello("Gladys")

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
	// message, err := greetings.Hello("Gladys")

	// Request greeting messages for the names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned print the returned message
	// to the console
	fmt.Println(messages)
}
