package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)
	go add(100, 200, ch)
	result := <-ch //RECEIVE
	fmt.Println("result :", result)
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result //SEND
}
