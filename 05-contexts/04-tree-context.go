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
	valCtx := context.WithValue(rootCtx, "k1", "v1")
	timeoutCtx, cancel := context.WithTimeout(valCtx, 10*time.Second)
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
	no := 1

	childWg := &sync.WaitGroup{}

	childWg.Add(1)
	evenValCtx := context.WithValue(stopCtx, "k1", nil)
	evenStopCtx, evenCancel := context.WithTimeout(evenValCtx, 3*time.Second)
	defer evenCancel()
	go printEvenNos(childWg, evenStopCtx)

	childWg.Add(1)
	oddStopCtx, oddCancel := context.WithTimeout(stopCtx, 7*time.Second)
	defer oddCancel()
	go printOddNos(childWg, oddStopCtx)
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
	childWg.Wait()
	fmt.Println("Finished generating nos")
}

func printEvenNos(wg *sync.WaitGroup, stopCtx context.Context) {
	defer wg.Done()
	fmt.Printf("[printEvenNos] value of k1 = %v\n", stopCtx.Value("k1"))
	no := 0
LOOP:
	for {
		select {
		case <-stopCtx.Done():
			break LOOP
		default:
			fmt.Println("[printEvenNos], no =", no)
			no = no + 2
			time.Sleep(300 * time.Millisecond)
		}

	}
	fmt.Println("Finished generating even nos")
}

func printOddNos(wg *sync.WaitGroup, stopCtx context.Context) {
	defer wg.Done()
	fmt.Printf("[printOddNos] value of k1 = %v\n", stopCtx.Value("k1"))
	no := 1
LOOP:
	for {
		select {
		case <-stopCtx.Done():
			break LOOP
		default:
			fmt.Println("[printOddNos], no =", no)
			no = no + 2
			time.Sleep(800 * time.Millisecond)
		}

	}
	fmt.Println("Finished generating odd nos")
}
