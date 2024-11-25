package main

import (
	"fmt"
	"gs/sort"
)

func main() {
	
	var array = []int{9,8,7,6,5,4,3,2,1}
	// sort.BubbleSort(array)
	new := sort.MergeSort(array)
	fmt.Println(new)
}
