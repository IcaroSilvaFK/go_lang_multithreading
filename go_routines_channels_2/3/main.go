package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go publish(ch)

	subscribe(ch)

}

func publish(c chan<- string) {

	for i := 0; i < 1_000; i++ {
		msg := fmt.Sprintf("Message %d", i)

		c <- msg
	}
	close(c) // indica que nada mais entra no canal
}

func subscribe(c <-chan string) {
	for msg := range c {
		fmt.Println(msg)
	}

}
