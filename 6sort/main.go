package main

import "fmt"

func main() {
	arr := []int{2, 1, 5, 23, 23, 100, 50}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	return
	mergesort(arr, 0, len(arr)-1)
	fmt.Println(midFind(arr, 100, 0, len(arr)-1))
}

func midFind(arr []int, l, r int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	if l >= r {
		return -1
	}

	if arr[(l+r)/2] == target {
		return (l + r) / 2
	}

	if arr[(l+r)/2] > target {
		return midFind(arr, l, (l+r)/2-1, target)
	} else {
		return midFind(arr, (l+r)/2+1, r, target)
	}
}

/**
l 左下标
r 右下标
*/
func firstMidFind(arr []int, l, r int, target int) int {
	low := l
	high := r - 1
	for low <= high {
		mid := low + (high-low)>>1

		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != 0 {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func lastMidFind(arr []int, l, r int, target int) int {
	low := l
	high := r - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == r-1 || arr[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func firstBiggerFind(arr []int, l, r int, target int) int {
	low := l
	high := r - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] >= target {
			if mid == 0 || arr[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

func firstLessFind(arr []int, l, r int, target int) int {
	low := l
	high := r - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] <= target {
			if mid == r-1 || arr[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
