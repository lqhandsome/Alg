package main

import "fmt"

func main() {
	//arr := []int{0, 3, 5, 1}
	arr := []int{5, 1, 1, 2, 0, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 快速排序
func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	left := l
	right := r
	// 主元
	pivot := arr[l]

	for l < r {

		// 从右开始处理数据
		for l < r && arr[r] >= pivot {
			r--
		}
		// 从左开始遍历
		for l < r && arr[l] <= pivot {
			l++
		}
		// 交换数据
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	// 交换主元
	arr[l], arr[left] = arr[left], arr[l]
	// 递归处理
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
}
