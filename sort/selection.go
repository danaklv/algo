package sort

func SelectionSort(arr []int) []int {
	var res []int
	
	for len(arr) >0 {
		smallest := findSmallest(arr)
		res = append(res, arr[smallest])
		arr = removeByIndex(arr, smallest)
	
	}

	return res

}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}

	}
	return smallestIndex
}

func removeByIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
