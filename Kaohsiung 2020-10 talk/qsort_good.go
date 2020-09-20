package qsort

/*
import "sync"

func qsortGoodWorker(inputCh chan []input, wg *sync.WaitGroup) {
	defer wg.Done()

	for input := range inputCh {
		if len(input) <= 1 {
			continue
		}

		pivotPos := qsortPartition(input)
		inputCh <- input[:pivotPos]
		inputCh <- input[pivotPos+1:]
	}
}

func qsortGood(input []int) {
	wg := sync.WaitGroup{}

}
*/
