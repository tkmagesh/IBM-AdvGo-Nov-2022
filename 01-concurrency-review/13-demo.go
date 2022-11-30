package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(4 * time.Second)
		d3 := <-ch3
		fmt.Println("ch3 :", d3)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	/*
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			fmt.Println("ch1 :", <-ch1)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			fmt.Println("ch2 :", <-ch2)
			wg.Done()
		}()
		wg.Wait()
	*/
	for i := 1; i <= 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println("ch1 :", d1)
		case ch3 <- 300:
			fmt.Println("Sent the data to ch3")
		case d2 := <-ch2:
			fmt.Println("ch2 :", d2)
		default:
			fmt.Println("No channel operations were successful")
		}
	}
}
