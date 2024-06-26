package graph

import log "github.com/sirupsen/logrus"

var found bool

// graph 无向无权图
type graph struct {
	v    uint
	data [][]int
}

func NewGraph(v uint) *graph {
	return &graph{
		v:    v,
		data: make([][]int, v+1, v+1),
	}
}

// Add 添加一个s->t的边
func (g *graph) Add(s, t int) error {
	g.data[s] = append(g.data[s], t)
	g.data[t] = append(g.data[t], s)
	return nil
}

// bfs
func (g *graph) bfs(s, t int) {
	if s < 0 {
		return
	}

	visited := make(map[int]bool, g.v)
	prev := make([]int, g.v+1)
	for key, _ := range prev {
		prev[key] = -1
	}

	var queue []int
	queue = append(queue, s)
	visited[s] = true
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		for _, value := range g.data[tmp] {
			if !visited[value] {
				prev[value] = tmp
				queue = append(queue, value)
			}

			if value == t {
				log.Info(prev)
				printPrev(s, t, prev)
				return
			}

			visited[value] = true

		}
	}
}

func (g *graph) dfs(s, t int) {
	if s == t {
		return
	}

	visited := make(map[int]bool, g.v)
	prev := make([]int, g.v+1)
	for key, _ := range prev {
		prev[key] = -1
	}
	visited[s] = true
	g.dfsHandle(visited, prev, s, s, t)
}

func (g *graph) dfsHandle(visited map[int]bool, prev []int, ss, s, t int) {
	if found {
		return
	}

	for _, value := range g.data[s] {
		if visited[value] {
			continue
		}

		prev[value] = s
		visited[value] = true
		if t == value {
			found = true
			printPrev(ss, t, prev)
		}

		g.dfsHandle(visited, prev, ss, value, t)
	}

}

func printPrev(s, t int, prev []int) {
	if t == s {
		log.Info(t, "->>")
		return
	}

	printPrev(s, prev[t], prev)
	log.Info(t, "->>")
}
