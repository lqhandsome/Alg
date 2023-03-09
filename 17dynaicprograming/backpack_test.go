package main

import "testing"

func BenchmarkBackpack(b *testing.B) {
	items := []int{2, 2, 4, 6, 3}
	mem = make([][]bool, len(items))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		backpack(0, 0, items, len(items), 100)
	}
}

func BenchmarkBackpackFilter(b *testing.B) {
	items := []int{2, 2, 4, 6, 3}
	mem = make([][]bool, len(items))
	for k, _ := range mem {
		mem[k] = make([]bool, 100)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		backpackFilter(0, 0, items, len(items), 100)
	}
}

func BenchmarkBackpackDya(b *testing.B) {
	items := []int{2, 2, 4, 6, 3}
	status := make([][]bool, len(items))
	for k, _ := range status {
		status[k] = make([]bool, 100+1)
	}
	for i := 0; i < b.N; i++ {
		backpackDya(status, items, len(items), 100)
	}
}

func BenchmarkEditDistanceHuiSu(b *testing.B) {
	st1 = "aaaaabcd"
	st2 = "abce"
	ijDist = make([][]int, len(st2))
	for k, _ := range ijDist {
		ijDist[k] = make([]int, len(st1))
	}
	for i := 0; i < b.N; i++ {
		lwstBT(st2, 0, len(st2), st1, 0, len(st1), 0)
	}
}

func BenchmarkEditDistanceHuiSuFilter(b *testing.B) {

	for i := 0; i < b.N; i++ {
		lwstBTNoFilter(st2, 0, len(st2), st1, 0, len(st1), 0)
	}
}

func BenchmarkEditDistanceHuiSuTwo(b *testing.B) {
	st1 = "aaaaabcd"
	st2 = "abce"

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ijDist = make([][]int, len(st2))
		for k, _ := range ijDist {
			ijDist[k] = make([]int, len(st1))
		}
		b.StartTimer()
		lwstBT(st2, 0, len(st2), st1, 0, len(st1), 0)
	}
}
