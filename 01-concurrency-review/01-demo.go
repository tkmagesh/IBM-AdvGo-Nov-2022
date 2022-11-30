package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //increment the counter by 1
	go f1()
	f2()
	wg.Wait() //wait for the counter to become 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
