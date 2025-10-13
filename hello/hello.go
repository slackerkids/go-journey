package main

import (
	"fmt"
	"log"

	"github.com/slackerkids/go-journey/greetings"
)

func main() {
	// Set properties of the predifined logger, including
	// the log entry prefix and a flag to disable priniting
	// the time, source file, and a line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
