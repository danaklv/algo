package main

import (
	"fmt"
	"gs/search"
	"gs/sort"
)

func main() {
	
	var array = []int{9,8,7,6,5,4,3,2,1}
	// sort.BubbleSort(array)
	new := sort.MergeSort(array)

	fmt.Println(new)
	fmt.Println(search.BinarySearch(new,6))
	
}
