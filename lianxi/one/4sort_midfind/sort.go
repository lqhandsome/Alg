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

// mergeSort 归并排序
func mergeSort(arr []uint) []uint {

	if len(arr) < 2 {
		return arr
	}
	arr = sort(arr)
	return arr
}

func sort(arr []uint) []uint {

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	a1 := sort(arr[0:mid])
	a2 := sort(arr[mid:])
	return mergeArr(a1, a2)
}

// mergeArr 合并两个有序数组
func mergeArr(arr1 []uint, arr2 []uint) []uint {

	var sumArr []uint
	for len(arr1) > 0 && len(arr2) > 0 {

		if arr1[0] < arr2[0] {
			sumArr = append(sumArr, arr1[0])
			arr1 = append(arr1[:0], arr1[1:]...)
		} else {
			sumArr = append(sumArr, arr2[0])
			arr2 = append(arr2[:0], arr2[1:]...)
		}
	}

	if len(arr1) > 0 {
		sumArr = append(sumArr, arr1...)
	}

	if len(arr2) > 0 {
		sumArr = append(sumArr, arr2...)
	}

	return sumArr
}
