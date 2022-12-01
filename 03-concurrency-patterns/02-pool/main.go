package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
	Resource = any object that implements io.Closer interface
	Factory = function that creates an instance of the resource
*/

//Resource
type DBConnection struct {
	Id int
}

func (dbc DBConnection) Close() error {
	fmt.Printf("Resource [id=%d] is being discarded\n", dbc.Id)
	return nil
}

//Factory

var idCounter int

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	dbConnection := DBConnection{
		Id: idCounter,
	}
	return dbConnection, nil
}

func main() {
	p, err := pool.New(5, DBConnectionFactory)
	if err != nil {
		log.Fatalln(err)
	}
	wg := &sync.WaitGroup{}
	clientCount := 20
	wg.Add(clientCount)
	for c := 1; c <= clientCount; c++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(c)
	}
	wg.Wait()
	fmt.Println("Batch - 1 completed.  Press ENTER to continue...")

	wg.Add(10)
	for c := 21; c <= 30; c++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(c)
	}
	wg.Wait()
	p.Close()
}

func doWork(client int, p pool.Pool) {
	connResource, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("worker [id=%d]: Acquired resource [id=%d]\n", client, connResource.(DBConnection).Id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) //simulation of using the resource
	fmt.Printf("worker [id=%d]: Releasing resource [id=%d]\n", client, connResource.(DBConnection).Id)
	p.Release(connResource)
}
