package main

var (
	MaxW int = -100
	mem  [][]bool
)

// 0-1背包问题
func init() {
	items := []int{2, 2, 4, 6, 3}
	mem = make([][]bool, len(items))
	for k, _ := range mem {
		mem[k] = make([]bool, 100)
	}
	backpack(0, 0, items, len(items), 100)
	//fmt.Println(MaxW)
}
func main() {
	//items := []int{2, 2, 4, 6, 3}
	//backpackDya(items, len(items), 100)
}

// 0-1背包问题
// cw 当前装进去的物品的重量和
// i 考察到哪个物品了
// items 物品集合
// n 背包中物品数量
// w 背包可以装的最大重量
func backpack(cw int, i int, items []int, n int, w int) {
	if cw == w || i == n {
		if cw > MaxW {
			MaxW = cw
		}
		return
	}
	// 不放
	backpack(cw, i+1, items, n, w)
	// 放
	if cw+items[i] <= w { // 修枝，如果要放入大于背包容量就没必要再放
		backpack(cw+items[i], i+1, items, n, w)
	}
}

// 不计算重复的节点
func backpackFilter(cw int, i int, items []int, n int, w int) {

	if cw == w || i == n {
		if cw > MaxW {
			MaxW = cw
		}
		return
	}
	if mem[i][cw] {
		return
	}
	mem[i][cw] = true
	// 不放
	backpackFilter(cw, i+1, items, n, w)
	// 放
	if cw+items[i] <= w { // 修枝，如果要放入大于背包容量就没必要再放
		backpackFilter(cw+items[i], i+1, items, n, w)
	}
}

// 动态规划
func backpackDya(status [][]bool, items []int, n int, w int) {
	//用一个二维数组记录状态集
	//status := make([][]bool, n)
	//for k, _ := range status {
	//	status[k] = make([]bool, w+1)
	//}
	status[0][2] = true
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { // 第i个物品不放入背包
			if status[i-1][j] { // 如果上一步也没有放入
				status[i][j] = status[i-1][j]
			}
		}

		for j := 0; j <= w-items[i]; j++ {
			if status[i-1][j] {
				status[i][j+items[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if status[n-1][i] {
			return
		}
	}
}
