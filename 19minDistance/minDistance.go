package main

import (
	"fmt"
	"log"
)

// Edge 有向有权图边
type Edge struct {
	S string // 起始点
	T string // 终点
	W uint   // 权重
}

type Graph struct {
	V    uint //顶点数
	Data map[string][]*Edge
}

func main() {
	graph := InitGraph(5)
	graph.AddEdge("s", "t", 1)
	fmt.Println(graph.Data)
}

func InitGraph(v uint) *Graph {
	return &Graph{
		v,
		make(map[string][]*Edge, v),
	}
}

// AddEdge 添加一条s->t的边；权重为w
func (p *Graph) AddEdge(s, t string, w uint) error {
	if p == nil {
		log.Println("graph is nil")
		return fmt.Errorf("%v", "graph is nil")
	}

	if s == "" || t == "" || w == 0 {
		log.Println("s,t,w is empty")
		return fmt.Errorf("%v", "s,t,w is empty")
	}

	p.Data[s] = append(p.Data[s], &Edge{
		s,
		t,
		w,
	})
	return nil
}
