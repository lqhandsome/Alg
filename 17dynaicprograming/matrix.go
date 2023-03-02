package main

import (
	"fmt"
	"math"
)

// 最短矩阵
// 从矩阵一点走到另一点的最短路径
var MinDist = math.MaxInt
var count int

type Matrix struct {
	Data   [][]int // 二维数组保存矩阵信息
	StartX int     // 起点X
	StartY int     // 起点Y坐标
	End    int     // 重点 x,y的坐标
	W      int     // 最小路径累计值
}

func InitMatrix(data [][]int, startX, startY int, end int) Matrix {

	return Matrix{
		Data:   data,
		StartX: startX,
		StartY: startY,
		End:    end,
		W:      math.MaxInt,
	}
}

func (p Matrix) Find(dist int, x, y int) {
	dist += p.Data[x][y]
	if x == p.End && y == p.End {
		count++
		if dist < MinDist {
			MinDist = dist
		}
		return
	}
	//向下走
	if x < p.End {
		p.Find(dist, x+1, y)
	}

	// 向右走
	if y < p.End {
		p.Find(dist, x, y+1)
	}
}

func main() {
	data := make([][]int, 4)
	data[0] = []int{1, 3, 5, 9}
	data[1] = []int{2, 1, 3, 4}
	data[2] = []int{5, 2, 6, 7}
	data[3] = []int{6, 8, 4, 3}
	matrix := InitMatrix(data, 0, 0, 3)
	matrix.Find(0, 0, 0)
	fmt.Println(MinDist, count)
}
