package main

import "fmt"

func main() {
	fmt.Println(fbnqRecursion(5))
	fmt.Println(fbnqRange(5))
}

// fbnqRecursion 斐波那契数列
func fbnqRecursion(n int) int {

	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fbnqRecursion(n-1) + fbnqRecursion(n-2)
}

// fbnqRange 斐波那契数列
func fbnqRange(n int) int {

	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	last := 1
	lastLast := 1

	sum := 0
	for n > 2 {
		sum = last + lastLast
		lastLast = last
		last = sum
		n--
	}
	return sum
}
