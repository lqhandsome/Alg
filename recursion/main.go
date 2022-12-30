package main

import "fmt"

func main() {
	fmt.Println(f(5))
}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}
