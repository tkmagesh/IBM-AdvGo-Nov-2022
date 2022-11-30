package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ch := genPrimes()
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes() <-chan int {
	ch := make(chan int)
	timeoutCh := timeout(5 * time.Second)
	no := 3
	go func() {
	LOOP:
		for {
			select {
			case <-timeoutCh:
				fmt.Println("Timeout!!")
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
