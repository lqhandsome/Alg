package main

import (
	"fmt"
)

var (
	count int
)

//func main() {
//	arr := make([]int, 8)
//	queue8(arr, 0)
//	fmt.Println(count)
//}

/*
row 行
*/
func queue8(arr []int, row int) {
	if row == 8 {
		print(arr)
		count++
		fmt.Println("-----------------")
		return
	}
	for i := 0; i < 8; i++ {
		//判断某一行某一列是否可以放入
		if isOk(arr, i, row) {
			arr[row] = i
			queue8(arr, row+1)
		}
	}
}

/*
*
col 列
row 行
*/
func isOk(arr []int, col, row int) bool {
	// 左侧
	leftUp := col - 1
	// 右侧
	rightUp := col + 1
	// 依次向上判断
	for i := row - 1; i >= 0; i-- {
		// 判断上方
		if arr[i] == col {
			return false
		}
		//判断左上角
		if leftUp >= 0 && arr[i] == leftUp {
			return false
		} else {
			leftUp--
		}
		//判断右上角
		if leftUp <= 7 && arr[i] == rightUp {
			return false
		} else {
			rightUp++
		}
	}
	return true
}
func print(queue []int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == queue[i] {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}
