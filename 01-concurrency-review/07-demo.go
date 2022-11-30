package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println("Result :", result)
}

func add(x, y int) int {
	fmt.Println("add started")
	time.Sleep(3 * time.Second)
	result := x + y
	fmt.Println("add completed")
	return result
}
