package main

import (
	"fmt"
	"time"
)

//consumer
func main() {

	ch := add(100, 200)
	/*
		go func() {
			ch <- 2000
		}()
	*/
	result := <-ch //RECEIVE
	fmt.Println("result :", result)
}

//producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result //SEND
	}()
	return ch
}
