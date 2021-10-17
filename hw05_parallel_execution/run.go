package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type counter struct {
	sync.RWMutex
	errCount int
}

func (c *counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.errCount++
}

func (c *counter) getCount() int {
	c.RLock()
	defer c.RUnlock()
	return c.errCount
}

func NewCounter() *counter {
	return new(counter)
}

var wg sync.WaitGroup
//var mu sync.Mutex

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	in := make(chan Task, len(tasks))
	out := make(chan struct{}, 1)
	counter := NewCounter()

	for _, task := range tasks {
		in <- task
	}
	close(in)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range in {
				if (*counter).getCount() >= m {
					out <- struct{}{}
					close(out)
				}
				err := task()

				if err != nil {
					(*counter).increment()
				}
				select {
				case <-out:
					return
				default:
					continue
				}
			}
		}()
	}
	wg.Wait()

	select {
	case <-out:
		if m <= 0 {
			return nil
		}
		return ErrErrorsLimitExceeded
	default:
		return nil
	}
}
