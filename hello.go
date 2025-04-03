package main

import (
	"fmt"
	"log"

	"example.com/greetingtest"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0) // do not show timestamp

	fmt.Println("Hello bich")
	fmt.Println(quote.Go())
	messages, err := greeting.Hellos([]string{"Jack", "Jill", "Joko"})
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
