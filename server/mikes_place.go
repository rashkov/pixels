package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	go func() {
		msg := <-messages
		fmt.Println(msg)
	}()

	// Pause for input
	var input string
	fmt.Scanln(&input)

}
