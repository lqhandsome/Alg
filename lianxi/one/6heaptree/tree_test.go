package main

import "testing"

func TestInit(t *testing.T) {

	tr := InitTree(7)
	tr.Add(6)
	tr.Add(8)
	tr.Add(8)
	tr.Add(8)
	tr.Add(5)
	tr.LevelRange()
}
