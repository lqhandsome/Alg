package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {

	g := NewGraph(10)
	g.add(0, 1)
	g.add(0, 3)
	g.add(1, 2)
	g.add(1, 4)
	g.add(2, 5)
	g.add(3, 4)
	g.add(4, 5)
	g.add(4, 6)
	g.add(5, 7)
	g.add(6, 7)
	g.dfs(1, 7)
}

func TestAAA(t *testing.T) {
	aaa()
}
func aaa() {
	a := []byte("aaa/bbb")
	index := bytes.Index(a, []byte("/"))
	b := a[:index]
	c := a[index+1:]
	println(b, c)
	b = append(b, "ccc"...)
	println(b, c)

	fmt.Println(index, string(c))
}
