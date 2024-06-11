package main

import "fmt"

type a = int

func main() {
	//graph := [...][]int{
	//	{1, 3},
	//	{0, 2, 4},
	//	{1, 5},
	//	{0, 4},
	//	{1, 3, 5, 6},
	//	{2, 4, 7},
	//	{4, 7},
	//	{5, 6},
	//}
	//fmt.Println(graph)
	g := InitGraph(10)
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)
	g.bfs(0, 7)
	fmt.Println("-----------")
	g.bfsRange(0)
	fmt.Println()
	fmt.Println("-----------")
	g.dfs(1, 7)
}

// 邻接矩阵方式表示图(无向图)
type Graph struct {
	V    uint          // 定点数
	Data map[int][]int // 保存数据
}

func InitGraph(v uint) *Graph {
	graph := &Graph{
		V:    v,
		Data: make(map[int][]int, v),
	}
	return graph
}

func (p *Graph) AddEdge(s, t int) error {
	if p == nil || p.Data == nil {
		panic("not init")
		return fmt.Errorf("not init")
	}
	p.Data[s] = append(p.Data[s], t)
	p.Data[t] = append(p.Data[t], s)
	return nil
}

// 广度优先遍历
func (p *Graph) bfsRange(t int) {
	var queue []int
	visited := make(map[int]bool, p.V)
	queue = append(queue, t)
	for len(queue) > 0 {
		tmpT := queue[0]
		if len(queue) > 0 {
			queue = queue[1:]
		}
		if visited[tmpT] {
			continue
		}
		visited[tmpT] = true
		fmt.Print(tmpT)
		for i := 0; i < len(p.Data[tmpT]); i++ {
			queue = append(queue, p.Data[tmpT][i])
		}
	}
}

// 广度优先搜索
func (p *Graph) bfs(t, s int) {
	if t == s {
		return
	}
	var queue []int
	queue = append(queue, t)
	// 记录节点是否被访问过
	visited := make(map[int]bool, p.V)
	visited[t] = true

	// prev 数组用来记录该节点的父节点
	var prev []int
	for i := 0; i < len(p.Data); i++ {
		prev = append(prev, -1)
	}

	for len(queue) > 0 {
		tmpT := queue[0]
		queue = queue[1:]
		for i := 0; i < len(p.Data[tmpT]); i++ {
			if visited[p.Data[tmpT][i]] {
				continue
			}
			prev[p.Data[tmpT][i]] = tmpT
			visited[p.Data[tmpT][i]] = true
			if p.Data[tmpT][i] == s {
				printPrev(prev, s)
				return
			}
			queue = append(queue, p.Data[tmpT][i])
		}
	}
}

// 打印路径
func printPrev(s []int, key int) {

	if s[key] == -1 {

		return
	}
	printPrev(s, s[key])
	fmt.Println(key)
}
