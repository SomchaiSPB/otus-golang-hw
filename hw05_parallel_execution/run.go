package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type counter struct {
	sync.Mutex
	errCount int
}

func (c *counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.errCount++
}

var wg sync.WaitGroup

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var in = make(chan Task, len(tasks))
	out := make(chan struct{}, 1)
	counter := new(counter)

	for _, task := range tasks {
		in <- task
	}
	close(in)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range in {
				if counter.errCount >= m {
					select {
					case out <- struct{}{}:
						close(out)
					}
				}
				err := task()
				if err != nil {
					counter.increment()
				}
				select {
				case <-out:
					return
				default:
					//
				}
			}
		}()
	}
	wg.Wait()

	select {
	case <-out:
		return ErrErrorsLimitExceeded
	default:
		return nil
	}
}
