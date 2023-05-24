package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Millisecond)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- 1
	}()

	select {
	case <-c1:
		fmt.Println("c1 received first")
	case <-c2:
		fmt.Println("c2 received first")
	}

}
