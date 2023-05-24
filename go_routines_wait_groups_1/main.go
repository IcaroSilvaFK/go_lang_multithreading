package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	// wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from first go routine")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from second go routine")

			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			fmt.Println("Hello from third go routine")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("Oi")
}
