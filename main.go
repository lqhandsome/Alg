package main

import (
	"fmt"
)

type Graph struct {
	V    uint                // 顶点数
	Data map[string][]string // 保存节点值
}

func InitGraph(v uint) *Graph {
	return &Graph{
		V:    v,
		Data: make(map[string][]string, v),
	}
}

// t依赖于s，添加一条s->t的边
func (p *Graph) AddEdge(s, t string) error {
	if p == nil {
		fmt.Println("graph not init")

		return fmt.Errorf("graph not init")
	}

	if s == "" || t == "" {
		fmt.Println("s or t is empty")

		return fmt.Errorf("s or t is empty")
	}

	p.Data[s] = append(p.Data[s], t)
	return nil
}

func (p *Graph) TopSort() error {
	if len(p.Data) == 0 {
		return nil
	}

	// 计算每个顶点的入度
	inEdge := make(map[string]uint, p.V)
	for k, v := range p.Data {
		if inEdge[k] == 0 {
			inEdge[k] = 0
		}

		for _, vv := range v {
			inEdge[vv]++
		}
	}

	fmt.Println("inEdge", inEdge)

	//模拟队列
	var edgeZeroQueue []string
	for k, v := range inEdge {
		if v == 0 {
			edgeZeroQueue = append(edgeZeroQueue, k)
		}
	}

	for len(edgeZeroQueue) > 0 {
		tmpEdgeZero := edgeZeroQueue[0]
		edgeZeroQueue = edgeZeroQueue[1:]

		// print
		fmt.Print(tmpEdgeZero, "->")
		// 将所有依赖tmpEdgeZero 的顶点-1
		for _, v := range p.Data[tmpEdgeZero] {
			inEdge[v]--
			if inEdge[v] == 0 {
				edgeZeroQueue = append(edgeZeroQueue, v)
			}
		}
	}
	return nil
}
func main() {
	graph := InitGraph(5)
	e := graph.AddEdge("a", "b")
	e = graph.AddEdge("b", "c")
	e = graph.AddEdge("c", "d")
	e = graph.AddEdge("e", "f")
	e = graph.AddEdge("d", "e")

	if e != nil {
		panic(e)
	}
	fmt.Println(graph.Data)
	graph.TopSort()
}
