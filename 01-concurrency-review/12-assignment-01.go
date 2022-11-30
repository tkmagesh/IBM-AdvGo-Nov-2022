package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ch := genPrimes(3, 100)
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
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
