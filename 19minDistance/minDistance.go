package main

// Edge 有向有权图边
type Edge struct {
	S string // 起始点
	E string // 终点
	W uint   // 权重
}

type Graph struct {
	V    uint //顶点数
	Data map[string]*Edge
}

func main() {

}

func InitGraph(v uint) *Graph {
	return &Graph{
		v,
		make(map[string]*Edge, v),
	}
}
