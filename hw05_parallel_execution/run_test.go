package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, int32(tasksCount), runTasksCount, "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestMyRun1(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("ignore errors", func(t *testing.T) {
		tasksCount := 77
		tasks := make([]Task, 0, tasksCount)

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				return fmt.Errorf("error from task %d", i)
			})
		}

		workersCount := 15
		maxErrorsCount := -0

		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)
	})
	t.Run("High tasks : fail in middle", func(t *testing.T) {
		tasksCount := 100000
		tasks := make([]Task, 0, tasksCount)
		var runTasksCount int32
		period := 100
		for i := 0; i < tasksCount; i++ {
			if i%period == 0 {
				tasks = append(tasks, func() error {
					atomic.AddInt32(&runTasksCount, 1)
					return fmt.Errorf("error from task %d", i)
				})
			} else {
				tasks = append(tasks, func() error {
					time.Sleep(time.Millisecond)
					atomic.AddInt32(&runTasksCount, 1)
					return nil
				})
			}
		}
		workersCount := 1000
		maxErrorsCount := 100
		err := Run(tasks, workersCount, maxErrorsCount)
		require.NotNil(t, err)
		maxCompletedAllowed := int32(period*maxErrorsCount + 1 + workersCount)
		require.LessOrEqual(t, runTasksCount, maxCompletedAllowed, "extra tasks were started")
	})
}
func TestMyRun2(t *testing.T) {
	t.Run("High tasks : valid end", func(t *testing.T) {
		tasksCount := 100000
		tasks := make([]Task, 0, tasksCount)
		var runTasksCount int32
		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}
		workersCount := 1000
		maxErrorsCount := 0
		err := Run(tasks, workersCount, maxErrorsCount)
		require.Nil(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
	})

	t.Run("Empty", func(t *testing.T) {
		tasks := make([]Task, 0)
		err := Run(tasks, 0, 0)
		require.Nil(t, err)
	})
}
