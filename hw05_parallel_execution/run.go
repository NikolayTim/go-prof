package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var errorsCount int32

type Task func() error

func worker(channel <-chan Task, results chan<- error, id, m int, wg *sync.WaitGroup) {
	for task := range channel {
		if atomic.LoadInt32(&errorsCount) >= int32(m) {
			break
		}

		result := task()
		if result != nil {
			atomic.AddInt32(&errorsCount, 1)
		}

		results <- result
	}

	wg.Done()
}

func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	defer resetErrorsCount()

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	channel := make(chan Task, len(tasks))
	results := make(chan error, len(tasks))

	if len(tasks) < n {
		n = len(tasks)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(channel, results, i, m, &wg)
	}

	for _, v := range tasks {
		channel <- v
	}

	close(channel)
	wg.Wait()
	close(results)

	if len(results) < len(tasks) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func resetErrorsCount() {
	atomic.StoreInt32(&errorsCount, 0)
}
