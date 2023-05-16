package main

func finMid(arr []uint, data uint, l, r int) int {

	if l > r {
		return -1
	}

	if l == r {
		if arr[l] == data {
			return l
		} else {
			return -1
		}
	}

	mid := (l-r)/2 + r
	pivot := arr[mid]
	if pivot == data {
		return mid
	} else if pivot > data {
		return finMid(arr, data, l, mid-1)
	} else {
		return finMid(arr, data, mid+1, r)
	}
	return -1
}
