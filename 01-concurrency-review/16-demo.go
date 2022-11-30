package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genPrimes(stopCh)
	fmt.Println("Hit ENTER to stop...")

	go func() {
		fmt.Scanln()
		//stopCh <- struct{}{}
		close(stopCh)
	}()

	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	no := 3
	go func() {
	LOOP:
		for {
			select {
			case <-stopCh:
				fmt.Println("stop signal received")
				break LOOP
			default:
				if isPrime(no) {
					ch <- no
					time.Sleep(500 * time.Millisecond)
				}
				no++
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= int(math.Sqrt(float64(no))); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
