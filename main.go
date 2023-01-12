package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		slc = []int{}
		n   = 100
		wg  sync.WaitGroup
	)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			slc = append(slc, i)
			println(slc)
			println(&slc)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for _, v := range slc {
		fmt.Println(v)
	}
	fmt.Println("done")
}
