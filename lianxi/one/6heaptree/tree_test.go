package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {

	tr := InitTree(7)
	tr.Add(6)
	tr.Add(8)
	tr.Add(8)
	tr.Add(8)
	tr.Add(5)
	tr.LevelRange()
}

func TestTreeSearch(t *testing.T) {
	tr := InitTree(7)
	tr.Add(6)
	tr.Add(8)
	tr.Add(8)
	tr.Add(8)
	tr.Add(5)
	cases := []struct {
		search int
		res    bool
	}{
		{1, false},
		{2, false},
		{5, true},
		{8, true},
		{7, true},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("%d", k+100), func(t *testing.T) {
			if v.res != tr.Search(v.search) {
				t.Errorf("error,key: %s", t.Name())
			}
		})
	}
}
