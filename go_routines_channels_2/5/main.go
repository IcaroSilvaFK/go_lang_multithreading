package main

import "fmt"

// receive only
func readerChan(name string, ch chan<- string) {
	ch <- name

}

// read only
func read(ch <-chan string) {

	fmt.Println(<-ch)

}

func main() {

	hello := make(chan string)

	go readerChan("hello", hello)
	read(hello)

}
