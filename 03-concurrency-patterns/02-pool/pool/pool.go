package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool closed")

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	sync.Mutex
}

func New(poolsize int, factory func() (io.Closer, error)) (*Pool, error) {
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolsize),
		closed:    false,
	}, nil
}

func (pool *Pool) Acquire() (io.Closer, error) {
	pool.Lock()
	defer pool.Unlock()

	select {
	case r, ok := <-pool.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return r, nil
	default:
		fmt.Println("Acquiring from the factory")
		return pool.factory()
	}
}

func (pool *Pool) Release(resource io.Closer) error {
	pool.Lock()
	defer pool.Unlock()
	select {
	case pool.resources <- resource:
		fmt.Println("Releasing the resource to the pool")
		return nil
	default:
		fmt.Println("Pool full. Discarding the resource")
		return resource.Close()
	}
}

func (pool *Pool) Close() {
	pool.Lock()
	defer pool.Unlock()
	if pool.closed {
		return
	}
	close(pool.resources)
	for resource := range pool.resources {
		resource.Close()
	}
}
