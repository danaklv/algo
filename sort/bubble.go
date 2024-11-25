package sort

func BubbleSort(array []int) {
	if len(array) <= 1 { 
		return
	}
	k := 1 
	flag := false

	for i := 0; i < len(array)-1; i++ { 
		for j := 0; j < len(array)-k; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j] 
				flag = true
			}

		}
		if !flag { //досрончый выход, если позиции ни разу не сменились, значит массив уже отсортирован 
			return
		}
		k++
	}

}
