package qsort

import (
	"math/rand"
	"testing"
)

func generateRandomSlice(n int) []int {

	slice := make([]int, n, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

func isAscSorted(slice []int) bool {

	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func TestQsortBad(t *testing.T) {
	array := generateRandomSlice(1000000)

	qsortBad(array)

	if isAscSorted(array) {
		t.Error("the sorting is buggy", array)
	}
}

func TestQsortGood(t *testing.T) {
	array := generateRandomSlice(1000000)

	qsortGood(array)

	if isAscSorted(array) {
		t.Error("the sorting is buggy", array)
	}
}
