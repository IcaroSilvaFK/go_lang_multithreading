package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1_000)
	go publish(ch)

	go subscribe(ch, &wg)

	wg.Wait()
}

func publish(c chan<- string) {

	for i := 0; i < 1_000; i++ {
		msg := fmt.Sprintf("Message %d", i)

		c <- msg
		time.Sleep(time.Second)
	}

	close(c)
}

func subscribe(c <-chan string, wg *sync.WaitGroup) {
	for msg := range c {
		fmt.Println(msg)
		wg.Done()
	}
}
