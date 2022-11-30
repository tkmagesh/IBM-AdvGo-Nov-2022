package main

import (
	"fmt"
)

/*
	Channel Behavior
	* A RECEIVE operation is ALWAYS a blocking operation
	* A SEND operation is BLOCKED until a RECEIVE operation is initiated

*/

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/

	ch := make(chan int)
	go func() {
		fmt.Println("Attempting to send the data")
		ch <- 100
		fmt.Println("Completed sending the data")
	}()
	data := <-ch
	fmt.Println(data)
}
