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
	stack := make(chan struct{}, n)
	var wg sync.WaitGroup
	ctr := 0

	for i := 0; i < len(tasks); i++ {
		mu.RLock()
		if ctr >= m {
			mu.RUnlock()
			break
		}
		mu.RUnlock()

		j := i
		stack <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-stack
				wg.Done()
			}()
			err := tasks[j]()
			if err != nil {
				mu.Lock()
				ctr++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	if ctr < m {
		return nil
	}
	return ErrErrorsLimitExceeded
}
