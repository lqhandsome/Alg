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
