package qsort

import "sync"

func qsortBadInternal(input []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// sentinal
	if len(input) <= 1 {
		return
	}

	pivotPos := qsortPartition(input)

	wg.Add(2)
	go qsortBadInternal(input[:pivotPos], wg)
	go qsortBadInternal(input[pivotPos+1:], wg)
}

func qsortBad(input []int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go qsortBadInternal(input, &wg)
	wg.Wait()
}
