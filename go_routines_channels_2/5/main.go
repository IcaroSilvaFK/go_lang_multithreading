package main

import "fmt"

func readerChan(name string, ch chan<- string) {
	ch <- name

}

func read(ch <-chan string) {

	fmt.Println(<-ch)

}

func main() {

	hello := make(chan string)

	go readerChan("hello", hello)
	read(hello)

}
