package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	stopCh := make(chan struct{})
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	wg.Add(1)
	go printNos(wg, stopCh)
	wg.Wait()
}

func printNos(wg *sync.WaitGroup, stopCh chan struct{}) {
	defer wg.Done()
	no := 1
LOOP:
	for {
		select {
		case <-stopCh:
			break LOOP
		default:
			fmt.Println("[printNos], no =", no)
			no++
			time.Sleep(500 * time.Millisecond)
		}

	}
	fmt.Println("Finished generating nos")
}
