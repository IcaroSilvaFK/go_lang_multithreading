package main

import (
	"fmt"
	"time"
)

func main() {

	var singleChan = make(chan int, 0)

	go func() {
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from first go routine")
			v := <-singleChan
			singleChan <- v + 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from second go routine")

			v := <-singleChan
			singleChan <- v + 1

			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from third go routine")
			v := <-singleChan
			singleChan <- v + 1
			time.Sleep(time.Second)
		}
	}()

	for v := range singleChan {
		fmt.Println(v)
	}
}
