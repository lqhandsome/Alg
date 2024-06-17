package sort

// popSort 冒泡排序
func popSort(data []int) {
	if len(data) < 2 {
		return
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// insertSort 插入排序
func insertSort(data []int) {
	if len(data) < 2 {
		return
	}
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		j := i
		for ; j > 0 && tmp < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		data[j] = tmp
	}

}

// quickSort 快速排序
func quickSort([]int) {

}

// selectSort 选择排序
func selectSort(data []int) {
	if len(data) < 2 {
		return
	}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
}
