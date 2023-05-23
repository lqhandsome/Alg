package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_forceSearch(t *testing.T) {
	s := "abc"
	p := "ba"
	log.Info(forceSearch(s, p))

	cases := []struct {
		s     string
		p     string
		index int
	}{
		{"a", "b", -1},
		{"a", "a", 0},
		{"ab", "a", 0},
		{"ab", "b", 1},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
			if index := forceSearch(v.s, v.p); index != v.index {
				t.Errorf("%v,index: %v", cases[k], index)
			}
		})
	}
}
