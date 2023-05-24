package main

import "fmt"

func main() {

	forever := make(chan string)

	fmt.Println("Hey")

	//deadlock
	<-forever

}
