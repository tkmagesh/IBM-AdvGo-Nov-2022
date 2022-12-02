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
	/*
		timeoutCtx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
		defer cancel()
	*/
	val1Ctx := context.WithValue(rootCtx, "k1", "v1")
	// val2Ctx := context.WithValue(val1Ctx, "k2", "v2")
	val2Ctx := context.WithValue(val1Ctx, "k1", "new-v1")
	timeoutCtx, cancel := context.WithTimeout(val2Ctx, 10*time.Second)
	defer cancel()
	go func() {
		fmt.Scanln()
		cancel()
	}()

	wg.Add(1)
	go printNos(wg, timeoutCtx)
	wg.Wait()
}

func printNos(wg *sync.WaitGroup, stopCtx context.Context) {
	defer wg.Done()
	fmt.Printf("[printNos] Data from context (k1): %v\n", stopCtx.Value("k1"))
	fmt.Printf("[printNos] Data from context (k2): %v\n", stopCtx.Value("k2"))
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
