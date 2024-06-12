package graph

import (
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
	g.bfs(1, 7)
}
