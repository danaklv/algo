package sort

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	index := rand.Intn(len(arr))
	pivot := arr[index]
	var less []int
	var greater []int
	var equals []int

	for i := range arr {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			greater = append(greater, arr[i])
		} else {
			equals = append(equals, arr[i])
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)

	res := append(append(less, equals...), greater...)
	

	return res
}
