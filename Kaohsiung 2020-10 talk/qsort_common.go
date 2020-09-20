package qsort

// take input[0] as pivot, and then run one iteration
func qsortPartition(input []int) (pivotPos int) {
	pivot := input[0]

	startIdx := 1
	endIdx := len(input) - 1

	for {
		// scan for the swapping pairs
		for startIdx < endIdx && input[startIdx] <= pivot {
			startIdx++
		}
		for startIdx < endIdx && input[endIdx] >= pivot {
			endIdx--
		}

		if startIdx >= endIdx {
			break
		}
		// perform swapping
		input[startIdx], input[endIdx] = input[endIdx], input[startIdx]
	}

	// put back the pivot into correct position
	input[0], input[startIdx] = input[startIdx], pivot
	return startIdx
}
