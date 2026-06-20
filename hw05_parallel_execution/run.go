package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.

func Run(tasks []Task, n, m int) error {
	// m<=0 : ignore errors
	if m <= 0 {
		m = len(tasks) + 1
	}

	var mu sync.RWMutex
	ctr := 0
	stack := make(chan Task, len(tasks))
	for _, task := range tasks {
		stack <- task
	}
	close(stack)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range stack {
				err := task()
				if err != nil {
					mu.Lock()
					ctr++
					mu.Unlock()
				}
				mu.RLock()
				if ctr >= m {
					mu.RUnlock()
					break
				}
				mu.RUnlock()
			}
		}()
	}

	wg.Wait()
	if ctr < m {
		return nil
	}
	return ErrErrorsLimitExceeded
}
