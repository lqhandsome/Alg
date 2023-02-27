package main

import "fmt"

var maxW int = 0 // 背包可以承受的最大重量

// 0-1背包问题
// cw 当前装进去的物品的重量和
// i 考察到哪个物品了
// items 物品集合
// n 背包中物品数量
// w 背包可以装的最大重量
func packsack(i int, cw int, items []int, n int, w int) {
	// 找到的已经满足背包或者全部装完了i == n
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	// 不放入背包
	packsack(i+1, cw, items, n, w)
	if cw+items[i] <= w { // 修枝，超过背包重量没必要再放
		//放入背包
		packsack(i+1, cw+items[i], items, n, w)
	}
}

func main() {

	items := []int{10, 8, 11, 2}
	packsack(0, 0, items, len(items), 12)
	fmt.Println(maxW)
}
