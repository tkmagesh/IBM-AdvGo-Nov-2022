package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	stopCtx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go printNos(wg, stopCtx)
	wg.Wait()
}

func printNos(wg *sync.WaitGroup, stopCtx context.Context) {
	defer wg.Done()
	no := 1
LOOP:
	for {
		select {
		case <-stopCtx.Done():
			break LOOP
		default:
			fmt.Println("[printNos], no =", no)
			no++
			time.Sleep(500 * time.Millisecond)
		}

	}
	fmt.Println("Finished generating nos")
}
