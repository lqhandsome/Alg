package graph

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
func (p *graph) Add(t, w int) error {

	p.data[t] = append(p.data[t], w)
	p.data[w] = append(p.data[w], t)
	return nil
}

// bfsRange 广度优先遍历
func (p *graph) bfsRange(t int) {

	// 记录节点是否访问过
	visited := make(map[int]bool, p.v)

	var queue []int
	queue = append(queue, t)
	for len(queue) > 0 {

	}
}
