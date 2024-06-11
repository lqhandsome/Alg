package main

import "fmt"

type Graph struct {
	data map[int][]int
	v    int
}

var found bool

func NewGraph(v int) *Graph {
	return &Graph{
		data: make(map[int][]int, v),
		v:    v,
	}
}

// add 无向图两天都需要添加一次
func (g *Graph) add(s, t int) {
	g.data[s] = append(g.data[s], t)
	g.data[t] = append(g.data[t], s)

	return
}

// dfs 深度优先搜索
func (g *Graph) dfs(s, t int) {
	if s == t {
		return
	}

	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}

	// 切片模拟队列
	var queue []int
	queue = append(queue, s)
	visited := make(map[int]bool, g.v)
	visited[s] = true
	g.ddfs(prev, queue, s, t, visited)
}

func (g *Graph) ddfs(prev []int, queue []int, s, t int, visited map[int]bool) {
	if s == t {
		return
	}
	if found {
		return
	}

	tmpV := queue[0]
	queue = queue[1:]
	for _, value := range g.data[tmpV] {
		if visited[value] {
			continue
		}

		visited[value] = true
		prev[value] = tmpV
		if value == t {
			print(prev, t)
			found = true
			return
		}

		queue = append(queue, value)
		g.ddfs(prev, queue, value, t, visited)
	}
}

// print打印路径
func print(prev []int, t int) {
	//if t <= 0 || t == -1 {
	//	return
	//}
	//
	//print(prev, prev[t])
	//fmt.Printf("%d--->", t)

	if prev[t] == -1 {

		return
	}
	print(prev, prev[t])
	fmt.Printf("%d--->", t)

}

// graph 无向无权图
type graph struct {
	v    int           // 定点数
	data map[int][]int // 储存数据
}

func New(v int) *graph {

	return &graph{
		v:    v,
		data: make(map[int][]int, v),
	}
}

// Add
// 添加一条t->w的边
//func (p *graph) Add(t, w int) error {
//
//	p.data[t] = append(p.data[t], w)
//	p.data[w] = append(p.data[w], t)
//	return nil
//}
//
//// bfsRange 广度优先遍历
//func (p *graph) bfsRange(t int) {
//
//	// 记录节点是否访问过
//	visited := make(map[int]bool, p.v)
//
//	var queue []int
//	queue = append(queue, t)
//	for len(queue) > 0 {
//
//	}
//}
