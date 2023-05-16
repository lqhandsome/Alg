package main

// PopSort 冒泡排序
func PopSort(arr []uint) {
	if len(arr) < 1 {
		return
	}

	for k, _ := range arr {
		for kk := 0; kk < len(arr)-k-1; kk++ {
			if arr[kk] > arr[kk+1] {
				arr[kk], arr[kk+1] = arr[kk+1], arr[kk]
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort(arr []uint) {
	if len(arr) < 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// InsertSort 插入排序
func InsertSort(arr []uint) {
	if len(arr) < 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		tag := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > tag; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tag
	}
}

// QuickSort 快排
func QuickSort(arr []uint, l, r int) {

	if l >= r {
		return
	}

	left := l
	right := r

	// 主元
	pivot := arr[l]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		// 从左边处理数据
		for left < right && arr[left] <= pivot {
			left++
		}

		// 交换位置
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	// 将主元换到中间
	arr[left], arr[l] = arr[l], arr[left]

	// 递归处理
	QuickSort(arr, l, left-1)
	QuickSort(arr, left+1, r)
}
