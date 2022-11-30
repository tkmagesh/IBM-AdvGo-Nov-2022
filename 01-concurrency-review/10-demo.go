package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, isChannelOpen := <-ch; isChannelOpen {
			fmt.Println(data)
			continue
		}
		fmt.Println("All the data received")
		break
	}
}

func genNos(ch chan int) {
	ch <- 10
	time.Sleep(500 * time.Millisecond)
	ch <- 20
	time.Sleep(500 * time.Millisecond)
	ch <- 30
	time.Sleep(500 * time.Millisecond)
	ch <- 40
	time.Sleep(500 * time.Millisecond)
	ch <- 50
	fmt.Println("No more data")
	close(ch)
}
