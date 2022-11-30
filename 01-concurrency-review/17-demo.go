package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 5)

	fmt.Println("Attempting to send the data")
	ch <- 100
	ch <- 200
	ch <- 300
	fmt.Println("Completed sending the data")

	fmt.Println(len(ch))
	data := <-ch
	fmt.Println(data)
	fmt.Println(len(ch))
}
