package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {

	var singleChan = make(chan string) // vazio
	// Thread 2
	go func() {

		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from first go routine")
			v := fmt.Sprintf("Hello from first go routine %d", i)
			singleChan <- v // cheio
			time.Sleep(time.Second)
		}

	}()

	for v := range singleChan { // esvazia
		fmt.Println(v)
	}

}
