package main

import (
	"fmt"
	"log"

	"example.com/greetingtest"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("HAHAHAHA: ")
	log.SetFlags(0) // do not show timestamp

	fmt.Println("Hello bich")
	fmt.Println(quote.Go())
	message, err := greeting.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
