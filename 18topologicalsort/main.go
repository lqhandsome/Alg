package main

import "fmt"

// 邻接矩阵方式表示图（有向图）
type Graph struct {
	V    uint                // 定点数
	Data map[string][]string // 保存数据
}

func InitGraph(v uint) *Graph {
	graph := &Graph{
		V:    v,
		Data: make(map[string][]string, v),
	}
	return graph
}

// s优先于t执行（t依赖于s）添加一条s->t的边，
func (p *Graph) AddEdge(s, t string) error {
	if p == nil {
		return fmt.Errorf("graph is nil")
	}
	p.Data[s] = append(p.Data[s], t)
	return nil
}

// 拓扑排序
func (p *Graph) topSortByKahn() {
	if len(p.Data) == 0 {
		fmt.Println("data is zero")
		return
	}
	// 计算入度
	capV := make(map[string]uint)

	for k, v := range p.Data {
		if capV[k] == 0 {
			capV[k] = 0
		}
		for _, vv := range v {
			capV[vv] = capV[vv] + 1
		}
	}

	// 切片模拟队列
	queue := make([]string, 0)
	for k, v := range capV {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	fmt.Println(capV, queue)

	for len(queue) > 0 {
		//弹出一个读入度为0的节点
		tmp := queue[0]
		queue = queue[1:]
		fmt.Print(tmp, "-->")
		for _, v := range p.Data[tmp] {

			capV[v] = capV[v] - 1
			if capV[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
}

func main() {

	graph := InitGraph(5)
	e := graph.AddEdge("a", "b")
	e = graph.AddEdge("a", "c")
	e = graph.AddEdge("d", "a")
	e = graph.AddEdge("d", "b")
	e = graph.AddEdge("e", "b")
	if e != nil {
		panic(e)
	}
	fmt.Println(graph.Data)
	graph.topSortByKahn()
}
