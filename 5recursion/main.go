package main

import "fmt"

func main() {
	fmt.Println(f(5))
	fmt.Println(ff(5))

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

func ff(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	ret := 0
	pre := 2
	prePre := 1
	for i := 3; i <= n; i++ {
		ret = pre + prePre
		prePre = pre
		pre = ret
	}
	return ret
}
