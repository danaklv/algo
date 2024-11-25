package sort


func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left := array[:len(array)/2] // левая часть массива
	right := array[len(array)/2:] // правая

	left = MergeSort(left) // рекурсия
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	
	leftIndex, rightIndex := 0,0

	for range len(left) + len(right) {
		if leftIndex < len(left) && rightIndex < len(right) {
			if left[leftIndex] <= right[rightIndex] {
				result = append(result, left[leftIndex])
				leftIndex++
			} else  {
				result = append(result, right[rightIndex])
				rightIndex ++
			}
		} else if leftIndex == len(left) { 
			result = append(result, right[rightIndex])
			rightIndex++
		} else if rightIndex == len(right) {
			result = append(result, left[leftIndex])
			leftIndex++
		}
	}

	return result



}