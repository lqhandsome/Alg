package main

import "fmt"

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	queue8(arr, 0, 0)
}

func queue8(arr [8]int, col, row int) {
	if col == 7 {
		print(arr)
		return
	}
	for i := row; i < 8; i++ {
		if isOk(arr, i, row) {
			arr[i] = i
		} else {

		}
	}
}

func isOk(arr [8]int, col, row int) bool {
	return true
}
func print(queue [8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j <= 8; j++ {
			if j == queue[i] {
				fmt.Printf("Q")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}
