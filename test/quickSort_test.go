package test

import (
	"gs/sort"
	"math/rand"
	
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	array := make([]int, 10)

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}

	expected := make([]int, 10)

	copy(expected, array)

	sort.BubbleSort(expected)

	quickSortRes := sort.QuickSort(array)

	i := 0
	for i != len(array) {
		if expected[i] != quickSortRes[i] {
			t.Errorf("Incorrect result, want %v, got %v", expected, quickSortRes)
			return
		}
		i++
	}

}
