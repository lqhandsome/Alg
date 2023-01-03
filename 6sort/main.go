package main

import "fmt"

func main() {
	arr := []int{2, 1, 5, 23, 23, 100, 50}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
	return
	mergesort(arr, 0, len(arr)-1)
	fmt.Println(midFind(arr, 100, 0, len(arr)-1))
}

func midFind(arr []int, tag int, l, r int) (index int) {

	if arr[(l+r)/2] == tag {

		return (l + r) / 2
	}
	if l == r && arr[l] != tag {
		return -1
	}
	if arr[(l+r)/2] > tag {
		return midFind(arr, tag, l, (l+r)/2)
	} else {
		return midFind(arr, tag, (l+r)/2+1, r)
	}

}


