package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	fileWg := &sync.WaitGroup{}
	fileWg.Add(1)
	go source("data1.dat", dataCh, fileWg)
	fileWg.Add(1)
	go source("data2.dat", dataCh, fileWg)

	processWg := &sync.WaitGroup{}
	evenCh, oddCh := splitter(dataCh, processWg)

	processWg.Add(1)
	evenSumCh := make(chan int)
	go sum(evenCh, evenSumCh, processWg)

	processWg.Add(1)
	oddSumCh := make(chan int)
	go sum(oddCh, oddSumCh, processWg)

	processWg.Add(1)
	go merger(evenSumCh, oddSumCh, processWg)

	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Done")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if val, err := strconv.Atoi(txt); err == nil {
			ch <- val
		}
	}
	fmt.Printf("data from %q has been processed\n", fileName)
}

func splitter(dataCh chan int, processWg *sync.WaitGroup) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	processWg.Add(1)
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		defer processWg.Done()
		for val := range dataCh {
			if val%2 == 0 {
				evenCh <- val
			} else {
				oddCh <- val
			}
		}
	}()
	return evenCh, oddCh
}

func sum(valCh <-chan int, valSumCh chan int, processWg *sync.WaitGroup) {
	total := 0
	for val := range valCh {
		total += val
	}
	valSumCh <- total
	processWg.Done()
}

func merger(evenSumCh, oddSumCh chan int, processWg *sync.WaitGroup) {
	defer processWg.Done()
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i < 2; i++ {
		select {
		case evenTotal := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even total : %d\n", evenTotal))
		case oddTotal := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd total : %d\n", oddTotal))
		}
	}
}
