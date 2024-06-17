package graph

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

var g *graph

func TestMain(m *testing.M) {
	g = NewGraph(7)
	g.Add(1, 2)
	g.Add(1, 3)
	g.Add(2, 4)
	g.Add(2, 5)
	g.Add(3, 4)
	g.Add(4, 6)
	g.Add(5, 6)
	g.Add(6, 7)
	log.Info(g.data)
	m.Run()
}

func Test_BFS(t *testing.T) {
	g.bfs(4, 1)
}

func Test_dfs(t *testing.T) {
	g.dfs(5, 7)
}

func TestQuickSort(t *testing.T) {
	sli := []int{2, 1, 0}
	quickSort(sli, 0, len(sli)-1)
}

func quickSort(data []int, l, r int) {
	if len(data) < 2 {
		return
	}

	if l > r {
		return
	}
	left := l
	right := r
	pivot := data[l] // 主元

	for left < right {
		for right > left && data[right] >= pivot {
			right--
		}

		for right > left && pivot >= data[left] {
			left++
		}

		if right > left {
			data[right], data[left] = data[left], data[right]
		}
	}
	fmt.Println(right, left, data)
	//交换主元
	data[l], data[left] = data[left], data[l]
	quickSort(data, l, left-1)
	quickSort(data, left+1, r)
}
