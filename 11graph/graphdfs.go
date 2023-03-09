package main

var found bool

// 深度优先搜索
func (p *Graph) dfs(s, t int) {
	visited := make(map[int]bool, p.V)
	prev := make([]int, p.V, p.V)
	for k, _ := range prev {
		prev[k] = -1
	}
	p.recurDfs(s, t, visited, prev)
	printPrev(prev, t)
}

func (p *Graph) recurDfs(s, t int, visited map[int]bool, prev []int) {
	if found == true {
		return
	}
	visited[s] = true
	if s == t {
		found = true
		return
	}
	for i := 0; i < len(p.Data[s]); i++ {
		tmp := p.Data[s][i]
		if !visited[tmp] {
			prev[tmp] = s
			p.recurDfs(tmp, t, visited, prev)
		}
	}
}
