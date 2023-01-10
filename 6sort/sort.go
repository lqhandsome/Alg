package main

import "fmt"

/*
插入排序
 */
func InsertSort() {
	arr := []int{2,1,4,3,100,3,5,100}
	for i := 1;i < len(arr);i++ {
		tmp := arr[i]
		j:=i
		for ; j > 0 &&tmp < arr[j-1] ;j-- {
			arr[j]= arr[j-1]
		}
		arr[j] = tmp
		fmt.Println(arr)
	}
	fmt.Println(arr)
}

//归并排序
func mergesort(arr []int, l, r int) {
	merge_sort_c(arr, l, r)
}

func merge_sort_c(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	merge_sort_c(arr, l, m)
	merge_sort_c(arr, m+1, r)

	tmp := []int{}
	i, j := l, m+1
	fmt.Println(i, j)
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	if i <= m {
		tmp = append(tmp, arr[i:m+1]...)
	}
	if j <= r {
		tmp = append(tmp, arr[j:r+1]...)
	}
	for i := 0; i < len(tmp); i++ {
		arr[l+i] = tmp[i]
	}
}

//快排
func QuickSort(arr []int,l,r int) {
	if l >= r {
		return
	}

	i,j := l,r
	pivot := arr[j]
	for i < j {

		for arr[i] <= pivot && i <j {
			i++
		}
		for arr[j] >= pivot &&  i< j{
			j--
		}
		// 交换
		if i <j {
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	//主元换到中间
	arr[r],arr[i] = arr[i],pivot
	//arr[i] = pivot

	//arr[i],arr[r] = pivot,arr[i]
fmt.Println(arr[i],i)
	QuickSort(arr,i+1,r)
	QuickSort(arr,l,i-1)
}

// 冒泡排序
func popSort(arr []int){
	if len(arr) <2 {
		return
	}
	for i:=0;i < len(arr);i++ {
		for j :=0;j < len(arr)-i-1;j++{
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

//选择排序
func selectSort(arr []int) {
	if len(arr) <2 {
		return
	}
	for i:=0;i < len(arr);i++ {
		for j :=i+1;j < len(arr);j++{
			if arr[i] > arr[j] {
				arr[j],arr[i] = arr[i],arr[j]
			}
		}
	}
}