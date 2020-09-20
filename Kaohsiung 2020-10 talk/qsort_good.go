package qsort

import (
	"runtime"
	"sync"
)

func qsortGoodWorker(inputCh chan []int, wg *sync.WaitGroup, remainingTaskNum *sync.WaitGroup) {
	defer wg.Done()

	for input := range inputCh {
		// end condition of recursion
		if len(input) <= 1 {
			remainingTaskNum.Done()
			continue
		}

		pivotPos := qsortPartition(input)

		// add the sub-tasks to the queue
		remainingTaskNum.Add(2)
		inputCh <- input[:pivotPos]
		inputCh <- input[pivotPos+1:]

		// mark the current task is done
		remainingTaskNum.Done()
	}
}

// WARNING: this qsortGood is for demo only, not for production usage.
// The actual performance of qsortGood is MUCH worse than the standard library
func qsortGood(input []int) {
	wg := sync.WaitGroup{}
	remainingTaskNum := sync.WaitGroup{}

	inputCh := make(chan []int, len(input)/2+1)
	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go qsortGoodWorker(inputCh, &wg, &remainingTaskNum)
	}

	// add the input to channel, and wait for all subtask completed
	remainingTaskNum.Add(1)
	inputCh <- input
	remainingTaskNum.Wait()

	// let worker thread die peacefully, we SHOULD NOT leave the worker thread behind
	close(inputCh)
	wg.Wait()
}
