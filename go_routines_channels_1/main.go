package main

import (
	"fmt"
	"time"
)

func main() {

	var singleChan = make(chan int, 0)

	// var wg sync.WaitGroup

	// wg.Add(3)

	go func() {
		// defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from first go routine")
			singleChan <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		// defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from second go routine")

			singleChan <- i

			time.Sleep(time.Second)
		}
	}()
	go func() {
		// defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from third go routine")
			singleChan <- i
			time.Sleep(time.Second)
		}
	}()

	for v := range singleChan {
		fmt.Println(v)
	}

	// wg.Wait()
	fmt.Println("Oi")
}
