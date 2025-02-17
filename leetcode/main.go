package main

import (
	"fmt"
	"sort"
)

func main() {
	aa := []int{1, 2, -3, -4, 2}
	fmt.Println(threeSum(aa))

	return
	//arr := []int{0, 3, 5, 1}
	arr := []int{5, 1, 1, 2, 0, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func threeSum(arr []int) [][]int {
	if len(arr) < 3 {
		return nil
	}
	var res [][]int
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		target := 0 - arr[i]
		three := len(arr) - 1
		for j := i + 1; j < len(arr); j++ {
			//排除相同的
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			for three > j && arr[three]+arr[j] > target {
				three--
			}
			if three <= j {
				break
			}
			if arr[three]+arr[j] == target {
				res = append(res, []int{arr[i], arr[j], arr[three]})
			}
		}
	}

	return res
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
